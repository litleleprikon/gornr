//------------------------------------------------------------------------
// Copyright(c) 2024 Litleleprikon VST.
//------------------------------------------------------------------------

#pragma once

#include "pluginterfaces/base/funknown.h"
#include "pluginterfaces/vst/vsttypes.h"

namespace Litleleprikon {
//------------------------------------------------------------------------
static const Steinberg::FUID kLitleleprikonDelayProcessorUID (0x97A9380F, 0x429C5302, 0xA5BEC26A, 0x884EF150);
static const Steinberg::FUID kLitleleprikonDelayControllerUID (0x0CB92F05, 0x3F8A5557, 0xBB33D05B, 0x67695067);

// parameter tags
enum {
    kDelayId = 100,
    kBypassId = 101
};

#define LitleleprikonDelayVST3Category "Fx"

//------------------------------------------------------------------------
} // namespace Litleleprikon
