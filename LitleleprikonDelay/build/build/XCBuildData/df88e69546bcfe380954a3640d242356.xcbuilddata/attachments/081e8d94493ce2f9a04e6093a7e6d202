#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/vst3sdk/public.sdk/samples/vst-hosting/validator
  codesign --force --verbose --sign - /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Debug/validator
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/vst3sdk/public.sdk/samples/vst-hosting/validator
  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Debug/validator -selftest
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/vst3sdk/public.sdk/samples/vst-hosting/validator
  codesign --force --verbose --sign - /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Release/validator
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/vst3sdk/public.sdk/samples/vst-hosting/validator
  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Release/validator -selftest
fi

