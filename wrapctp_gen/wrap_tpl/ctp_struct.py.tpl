#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
# LUX et VERITAS
# Create On: 2025/04/06 20:37:58


from ctypes import Structure
from pyctp.[[ .Pf ]].ctp_datatype import *

[[ range .St ]]
class  [[ .FuncTypeName ]](Structure):
    """[[ .Comment ]]"""
    _fields_ = [
        [[ range .FuncFields ]]("[[ .FieldName ]]", [[ .FieldType|baseType ]]),
        [[ end ]]
    ]
    [[ end ]]