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

package GOLDILOCKS

var Modulus= [...]Chunk {0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FDFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFF}
const MConst Chunk=0x1

const CURVE_A int=1;
var CURVE_B = [...]Chunk {0x1FFF6756,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FDFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFF};
var CURVE_Order=[...]Chunk {0xB5844F3,0x1BC61495,0x1163D548,0x1984E51B,0x3690216,0xDA4D76B,0xFA7113B,0x1FEF9944,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x7FF};
var CURVE_Gx =[...]Chunk {0x15555555,0xAAAAAAA,0x15555555,0xAAAAAAA,0x15555555,0xAAAAAAA,0x15555555,0x152AAAAA,0xAAAAAAA,0x15555555,0xAAAAAAA,0x15555555,0xAAAAAAA,0x15555555,0xAAAAAAA,0x1555};
var CURVE_Gy =[...]Chunk {0xA9386ED,0x1757DE6F,0x13681AF6,0x19657DA3,0x3098BBB,0x12C19D15,0x12E03595,0xE515B18,0x17B7E36D,0x1AC426E,0xDBB5E8,0x10D8560,0x159D6205,0xB8246D9,0x17A58D2B,0x15C0};
