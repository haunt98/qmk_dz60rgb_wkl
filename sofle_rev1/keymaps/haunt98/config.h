// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_debounce_type.md
#ifdef DEBOUNCE
#    undef DEBOUNCE
#endif
#define DEBOUNCE 5

// https://github.com/qmk/qmk_firmware/blob/master/docs/tap_hold.md
#undef TAPPING_TERM
#define TAPPING_TERM 200
#undef QUICK_TAP_TERM
#define QUICK_TAP_TERM 120

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_encoders.md
#ifdef ENCODER_RESOLUTION
#    undef ENCODER_RESOLUTION
#endif
#define ENCODER_RESOLUTION 4
