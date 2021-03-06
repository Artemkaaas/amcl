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

package MF254M

var Modulus= [...]Chunk {0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0x3F80FFFF}
const MConst Chunk=0x3F810000

const CURVE_A int= -55790
var CURVE_B = [...]Chunk {0x0,0x0,0x0,0x0,0x0} // not used
var CURVE_Order=[...]Chunk {0xF3D3FEC46E98C7,0x306C8BD62FB0EA,0xFFFFFFFFFFEB95,0xFFFFFFFFFFFFFF,0xFE03FFF}
var CURVE_Gx =[...]Chunk {0x3,0x0,0x0,0x0,0x0}
var CURVE_Gy =[...]Chunk {0x0,0x0,0x0,0x0,0x0} // not used
