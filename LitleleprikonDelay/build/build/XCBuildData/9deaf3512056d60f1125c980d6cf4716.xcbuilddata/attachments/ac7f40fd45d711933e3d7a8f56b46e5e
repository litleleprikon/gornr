#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build
  codesign --force --verbose --sign - /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Debug/LitleleprikonDelay.vst3
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin
  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Debug/moduleinfotool -create -version 1.0.0.0 -path /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Debug/LitleleprikonDelay.vst3 -output /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Debug/LitleleprikonDelay.vst3/Contents/Resources/moduleinfo.json
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin
  codesign -f -s - -v /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Debug/LitleleprikonDelay.vst3/Contents/Resources/moduleinfo.json
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin
  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Debug/validator /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Debug/LitleleprikonDelay.vst3 
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build
  /opt/homebrew/Cellar/cmake/3.29.3/bin/cmake -E make_directory /Users/litleleprikon/Library/Audio/Plug-Ins/VST3
  ln -svfF /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Debug/LitleleprikonDelay.vst3 /Users/litleleprikon/Library/Audio/Plug-Ins/VST3
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build
  codesign --force --verbose --sign - /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Release/LitleleprikonDelay.vst3
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin
  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Release/moduleinfotool -create -version 1.0.0.0 -path /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Release/LitleleprikonDelay.vst3 -output /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Release/LitleleprikonDelay.vst3/Contents/Resources/moduleinfo.json
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin
  codesign -f -s - -v /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Release/LitleleprikonDelay.vst3/Contents/Resources/moduleinfo.json
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin
  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/bin/Release/validator  /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Release/LitleleprikonDelay.vst3
  cd /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build
  /opt/homebrew/Cellar/cmake/3.29.3/bin/cmake -E make_directory /Users/litleleprikon/Library/Audio/Plug-Ins/VST3
  ln -svfF /Users/litleleprikon/Documents/projects/github.com/litleleprikon/gornr/LitleleprikonDelay/build/VST3/Release/LitleleprikonDelay.vst3 /Users/litleleprikon/Library/Audio/Plug-Ins/VST3
fi

