cmake_minimum_required(VERSION 3.14.0)

option(SMTG_ENABLE_VST3_PLUGIN_EXAMPLES "Enable VST 3 Plug-in Examples" OFF)
option(SMTG_ENABLE_VST3_HOSTING_EXAMPLES "Enable VST 3 Hosting Examples" OFF)

set(CMAKE_OSX_DEPLOYMENT_TARGET 10.13 CACHE STRING "")

set(vst3sdk_SOURCE_DIR "/Users/litleleprikon/Documents/projects/github.com/steinbergmedia/vst3sdk")
if(NOT vst3sdk_SOURCE_DIR)
    message(FATAL_ERROR "Path to VST3 SDK is empty!")
endif()

project(LitleleprikonDelay
    # This is your plug-in version number. Change it here only.
    # Version number symbols usable in C++ can be found in
    # source/version.h and ${PROJECT_BINARY_DIR}/projectversion.h.
    VERSION 1.0.0.0 
    DESCRIPTION "LitleleprikonDelay VST 3 Plug-in"
)

set(SMTG_VSTGUI_ROOT "${vst3sdk_SOURCE_DIR}")

add_subdirectory(${vst3sdk_SOURCE_DIR} ${PROJECT_BINARY_DIR}/vst3sdk)
smtg_enable_vst3_sdk()

smtg_add_vst3plugin(LitleleprikonDelay
    source/version.h
    source/cids.h
    source/processor.h
    source/processor.cpp
    source/controller.h
    source/controller.cpp
    source/entry.cpp
    source/goprocess.h
)

#- VSTGUI Wanted ----
if(SMTG_ENABLE_VSTGUI_SUPPORT)
    target_sources(LitleleprikonDelay
        PRIVATE
            resource/editor.uidesc
    )
    target_link_libraries(LitleleprikonDelay
        PRIVATE
            vstgui_support
    )
    smtg_target_add_plugin_resources(LitleleprikonDelay
        RESOURCES
            "resource/editor.uidesc"
    )
endif(SMTG_ENABLE_VSTGUI_SUPPORT)
# -------------------

smtg_target_add_plugin_snapshots (LitleleprikonDelay
    RESOURCES
        resource/97A9380F429C5302A5BEC26A884EF150_snapshot.png
        resource/97A9380F429C5302A5BEC26A884EF150_snapshot_2.0x.png
)

target_link_libraries(LitleleprikonDelay
    PRIVATE
        sdk
)

smtg_target_configure_version_file(LitleleprikonDelay)

if(SMTG_MAC)
    smtg_target_set_bundle(LitleleprikonDelay
        BUNDLE_IDENTIFIER delay.vst.dev.litleleprikon
        COMPANY_NAME "Litleleprikon VST"
    )
    smtg_target_set_debug_executable(LitleleprikonDelay
        "/Applications/VST3PluginTestHost.app"
        "--pluginfolder;$(BUILT_PRODUCTS_DIR)"
    )
elseif(SMTG_WIN)
    target_sources(LitleleprikonDelay PRIVATE 
        resource/win32resource.rc
    )
    if(MSVC)
        set_property(DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR} PROPERTY VS_STARTUP_PROJECT LitleleprikonDelay)

        smtg_target_set_debug_executable(LitleleprikonDelay
            "$(ProgramW6432)/Steinberg/VST3PluginTestHost/VST3PluginTestHost.exe"
            "--pluginfolder \"$(OutDir)/\""
        )
    endif()
endif(SMTG_MAC)
