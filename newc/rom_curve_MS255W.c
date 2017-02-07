#include "arch.h"
#include "ecp_MS255W.h"

/* MS255 NUMS Curve - Weierstrass */

/* Pseudo-Mersenne NUMS curves http://eprint.iacr.org/2014/130 */

#if CHUNK==16
const int CURVE_A_MS255W=-3;
const BIG_256 CURVE_Order_MS255W= {0xAEB,0x2CA,0x60F,0x1F59,0x1AB8,0x19EF,0xB4A,0x507,0xA38,0x1C32,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0xFF};
const BIG_256 CURVE_B_MS255W= {0xB46,0x1FFD,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0x1FFF,0xFF};
const BIG_256 CURVE_Gx_MS255W= {0x1};
const BIG_256 CURVE_Gy_MS255W= {0x4BA,0xE5A,0x19DA,0x7FE,0x1173,0x1306,0x17F6,0x119B,0x16C7,0x1BED,0x1F45,0x5FE,0x1F9B,0xF5F,0xE48,0xF06,0xDBA,0x1607,0x1E9A,0xDE};

#endif

#if CHUNK==32
const int CURVE_A_MS255W=-3;
const BIG_256 CURVE_Order_MS255W= {0x1C594AEB,0x1C7D64C1,0x14ACF7EA,0x14705075,0x1FFFF864,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x7FFFFF};
const BIG_256 CURVE_B_MS255W= {0x1FFFAB46,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x1FFFFFFF,0x7FFFFF};
const BIG_256 CURVE_Gx_MS255W= {0x1};
const BIG_256 CURVE_Gy_MS255W= {0x9CB44BA,0x199FFB3B,0x1F698345,0xD8F19BB,0x17D177DB,0x1FFCD97F,0xCE487A,0x181DB74F,0x6F7A6A};
#endif

#if CHUNK==64
const int CURVE_A_MS255W=-3;
const BIG_256 CURVE_Order_MS255W= {0x8FAC983C594AEB,0x38283AD2B3DFAB,0xFFFFFFFFFF864A,0xFFFFFFFFFFFFFF,0x7FFFFFFF};
const BIG_256 CURVE_B_MS255W= {0xFFFFFFFFFFAB46,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0xFFFFFFFFFFFFFF,0x7FFFFFFF};
const BIG_256 CURVE_Gx_MS255W= {0x1};
const BIG_256 CURVE_Gy_MS255W= {0x33FF6769CB44BA,0xC78CDDFDA60D17,0xF9B2FF7D177DB6,0xEDBA7833921EBF,0x6F7A6AC0};
#endif