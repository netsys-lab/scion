#!/usr/bin/python3
# Copyright 2016 ETH Zurich
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""
:mod:`scmp_error_test` --- SCMP error tests
===========================================
"""
# Stdlib
import copy
import logging
import random
import sys

# SCION
from lib.defines import MAX_HOPBYHOP_EXT
from lib.main import main_wrapper
from lib.packet.ext.traceroute import TracerouteExt
from lib.packet.host_addr import HostAddrSVC
from lib.packet.ifid import IFIDPayload
from lib.packet.path import SCIONPath
from lib.packet.scion_addr import ISD_AS, SCIONAddr
from lib.packet.scmp.ext import SCMPExt
from lib.packet.scmp.types import (
    SCMPClass,
    SCMPCmnHdrClass,
    SCMPExtClass,
    SCMPPathClass,
    SCMPRoutingClass,
)
from lib.types import L4Proto
from test.integration.base_cli_srv import (
    TestClientBase,
    TestClientServerBase,
    setup_main,
)


class ErrorGenBase(TestClientBase):
    CLASS = None
    TYPE = None

    def run(self):
        # FIXME(kormat): only useful for errors generated by routers.
        if self._should_skip():
            return True
        self._send()
        if self.CLASS is None:
            # Allow testing of errors which don't send SCMP responses.
            return True
        pkt = self._recv()
        if not pkt:
            logging.error("Test timed out")
            return False
        ret = self._handle_response(pkt)
        self._shutdown()
        return ret

    def _should_skip(self):
        if self.addr.isd_as == self.dst.isd_as:
            return True

    def _send_raw_pkt(self, packed, next_hop, port):
        self.sock.send(packed, (str(next_hop), port))

    def _handle_response(self, spkt):
        spkt.parse_payload()
        l4 = spkt.l4_hdr
        if (l4.TYPE == L4Proto.SCMP and l4.class_ == self.CLASS and
                l4.type == self.TYPE):
            logging.debug("Success!\n%s", spkt)
            return True
        logging.error("Failure:\n%s", spkt)
        return False


# FIXME(kormat): ignore this for now, as with UDP-overlay, many packets are
# actually over MTU, so this check is disabled until we get TCP/SCION.
# class ErrorGenOversizePkt(ErrorGenBase):
#    CLASS = SCMPClass.ROUTING
#    TYPE = SCMPRoutingClass.OVERSIZE_PKT
#    DESC = "oversized packet"
#
#    def _create_payload(self, spkt):
#        padding = self.path.mtu - len(spkt) - len(self.data) + 1
#        return PayloadRaw(self.data + bytes(padding))


class ErrorGenBadHost(ErrorGenBase):
    CLASS = SCMPClass.ROUTING
    TYPE = SCMPRoutingClass.BAD_HOST
    DESC = "bad host"

    def _build_pkt(self):
        pkt = super()._build_pkt()
        pkt.set_payload(IFIDPayload.from_values(77))
        pkt.addrs.dst.host = HostAddrSVC(99, raw=False)
        return pkt


class ErrorGenBadPktLenShort(ErrorGenBase):
    CLASS = SCMPClass.CMNHDR
    TYPE = SCMPCmnHdrClass.BAD_PKT_LEN
    DESC = "bad pkt length (data missing)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        self._send_raw_pkt(spkt.pack()[:-1], next_hop, port)


class ErrorGenBadPktLenLong(ErrorGenBase):
    CLASS = SCMPClass.CMNHDR
    TYPE = SCMPCmnHdrClass.BAD_PKT_LEN
    DESC = "bad pkt length (extra data)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        self._send_raw_pkt(spkt.pack() + bytes([0]), next_hop, port)


class ErrorGenBadHdrLenShort(ErrorGenBase):
    DESC = "bad hdr length (too short)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        barr = bytearray(spkt.pack())
        barr[7] = 5
        self._send_raw_pkt(barr, next_hop, port)


class ErrorGenBadHdrLenLong(ErrorGenBase):
    DESC = "bad hdr length (too long)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        barr = bytearray(spkt.pack())
        barr[7] = 255
        self._send_raw_pkt(barr, next_hop, port)


class ErrorGenBadIOFOffsetShort(ErrorGenBase):
    CLASS = SCMPClass.CMNHDR
    TYPE = SCMPCmnHdrClass.BAD_IOF_OFFSET
    DESC = "bad IOF offset (too short)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        barr = bytearray(spkt.pack())
        barr[4] -= 8
        self._send_raw_pkt(barr, next_hop, port)


class ErrorGenBadIOFOffsetLong(ErrorGenBase):
    CLASS = SCMPClass.CMNHDR
    TYPE = SCMPCmnHdrClass.BAD_IOF_OFFSET
    DESC = "bad IOF offset (too long)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        barr = bytearray(spkt.pack())
        barr[4] = 255
        self._send_raw_pkt(barr, next_hop, port)


class ErrorGenBadHOFOffsetShort(ErrorGenBase):
    CLASS = SCMPClass.CMNHDR
    TYPE = SCMPCmnHdrClass.BAD_HOF_OFFSET
    DESC = "bad HOF offset (too short)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        barr = bytearray(spkt.pack())
        barr[5] = 1
        self._send_raw_pkt(barr, next_hop, port)


class ErrorGenBadHOFOffsetLong(ErrorGenBase):
    CLASS = SCMPClass.CMNHDR
    TYPE = SCMPCmnHdrClass.BAD_HOF_OFFSET
    DESC = "bad HOF offset (too long)"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        barr = bytearray(spkt.pack())
        barr[5] = 255
        self._send_raw_pkt(barr, next_hop, port)


class ErrorGenPathReq(ErrorGenBase):
    CLASS = SCMPClass.PATH
    TYPE = SCMPPathClass.PATH_REQUIRED
    DESC = "path required"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        spkt.path = SCIONPath()
        self._send_raw_pkt(spkt.pack(), next_hop, port)


class ErrorGenBadMAC(ErrorGenBase):
    CLASS = SCMPClass.PATH
    TYPE = SCMPPathClass.BAD_MAC
    DESC = "Bad MAC"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        hof = spkt.path.get_hof()
        hof.mac = bytes(hof.MAC_LEN)
        self._send_raw_pkt(spkt.pack(), next_hop, port)


class ErrorGenExpiredHOF(ErrorGenBase):
    CLASS = SCMPClass.PATH
    TYPE = SCMPPathClass.EXPIRED_HOF
    DESC = "Expired HOF"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        iof = spkt.path.get_iof()
        iof.timestamp = 0
        self._send_raw_pkt(spkt.pack(), next_hop, port)


class ErrorGenBadIF(ErrorGenBase):
    CLASS = SCMPClass.PATH
    TYPE = SCMPPathClass.BAD_IF
    DESC = "Bad Interface"

    def _should_skip(self):
        if super()._should_skip():
            return True
        # Only run if the remote AS is a core AS.
        if not self.sd.is_core_as(self.dst.isd_as):
            logging.info("Skipping non-core remote AS")
            return True
        return False

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        spkt.addrs.dst.isd_as = ISD_AS.from_values(3, 33)
        self._send_raw_pkt(spkt.pack(), next_hop, port)


class ErrorGenNonRoutingHOF(ErrorGenBase):
    CLASS = SCMPClass.PATH
    TYPE = SCMPPathClass.NON_ROUTING_HOF
    DESC = "Non-routing HOF"

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        hof = spkt.path.get_hof()
        hof.verify_only = True
        self._send_raw_pkt(spkt.pack(), next_hop, port)


class ErrorGenTooManyHbH(ErrorGenBase):
    CLASS = SCMPClass.EXT
    TYPE = SCMPExtClass.TOO_MANY_HOPBYHOP
    DESC = "Too many hop-by-hop extensions"

    def _create_extensions(self):
        exts = []
        for i in range(MAX_HOPBYHOP_EXT + 1):
            exts.append(TracerouteExt.from_values(5))
        return exts


class ErrorGenBadExtOrder(ErrorGenBase):
    CLASS = SCMPClass.EXT
    TYPE = SCMPExtClass.BAD_EXT_ORDER
    DESC = "Bad extension order"

    def _create_extensions(self):
        exts = []
        exts.append(TracerouteExt.from_values(5))
        exts.append(SCMPExt())
        return exts


class ErrorGenBadHopByHop(ErrorGenBase):
    CLASS = SCMPClass.EXT
    TYPE = SCMPExtClass.BAD_HOPBYHOP
    DESC = "Bad hop-by-hop extension"

    def _create_extensions(self):
        return [TracerouteExt.from_values(5)]

    def _send_pkt(self, spkt):
        next_hop, port = self.sd.get_first_hop(spkt)
        spkt.update()
        barr = bytearray(spkt.pack())
        idx = spkt.cmn_hdr.hdr_len + 2
        barr[idx] = 255
        self._send_raw_pkt(barr, next_hop, port)


GEN_LIST = (
    # ErrorGenOversizePkt,
    ErrorGenBadHost,
    ErrorGenBadPktLenShort,
    ErrorGenBadPktLenLong,
    ErrorGenBadHdrLenShort,
    ErrorGenBadHdrLenLong,
    ErrorGenBadIOFOffsetShort,
    ErrorGenBadIOFOffsetLong,
    ErrorGenBadHOFOffsetShort,
    ErrorGenBadHOFOffsetLong,
    ErrorGenPathReq,
    ErrorGenBadMAC,
    ErrorGenExpiredHOF,
    ErrorGenBadIF,
    ErrorGenNonRoutingHOF,
    ErrorGenTooManyHbH,
    ErrorGenBadExtOrder,
    ErrorGenBadHopByHop,
)


class SCMPErrorTest(TestClientServerBase):
    def __init__(self, client, server, sources, destinations):
        super().__init__(client, server, sources, destinations, local=False)
        self.src = client
        self.dst = server
        self.thread_name = "SCMPErr.MainThread"

    def _run(self):
        for src_ia in self.src_ias:
            src = SCIONAddr.from_values(src_ia, self.client_ip)
            # Pick N random non-local ASes as destination.
            for dst in self.iter_dsts(src_ia, 3):
                self._run_test(src, dst)

    def iter_dsts(self, src_ia, count):
        dst_ias = self.dst_ias[:]
        if src_ia in dst_ias:
            dst_ias.remove(src_ia)
        random.shuffle(dst_ias)
        for dst_ia in dst_ias[:count]:
            yield SCIONAddr.from_values(dst_ia, self.server_ip)

    def _run_test(self, src, dst):
        logging.info("=======================> Testing: %s -> %s",
                     src.isd_as, dst.isd_as)
        data = ("%s<->%s" % (src, dst)).encode("UTF-8")
        for cls_ in GEN_LIST:
            logging.info("===========> Testing: %s", cls_.DESC)
            client = cls_(self._run_sciond(src), copy.deepcopy(data), None,
                          copy.deepcopy(src), copy.deepcopy(dst), 0, api=True)
            if not client.run():
                sys.exit(1)

    def _create_data(self, src, dst):
        return ("%s<->%s" % (self.src, self.dst)).encode("UTF-8")


def main():
    args, srcs, dsts = setup_main("scmp_error_test")
    SCMPErrorTest(args.client, args.server, srcs, dsts).run()


if __name__ == "__main__":
    main_wrapper(main)
