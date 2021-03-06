/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

/* Fixed Data in ROM - Field and Curve parameters */

package MS256E

var Modulus= [...]Chunk {0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFA7FFFF}
const MConst Chunk=0xFFA80000

const CURVE_A int= -1
var CURVE_B = [...]Chunk{0x3BEE,0x0,0x0,0x0,0x0}
var CURVE_Order=[...]Chunk{0xB84E6F1122B4AD,0xA55AD0A6BC64E5,0xFFFFFFFFFFBE6A,0xFFFFFFFFFFFFFF,0x3FFFFFFF}
var CURVE_Gx =[...]Chunk{0xD,0x0,0x0,0x0,0x0}
var CURVE_Gy =[...]Chunk{0x7F6FB5331CADBA,0x6D63824D303F70,0xB39FA046BFBE2A,0x2A1276DBA3D330,0x7D0AB41E}

