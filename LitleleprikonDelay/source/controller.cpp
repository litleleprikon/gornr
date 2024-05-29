//------------------------------------------------------------------------
// Copyright(c) 2024 Litleleprikon VST.
//------------------------------------------------------------------------

#include "controller.h"
#include "cids.h"


using namespace Steinberg;

namespace Litleleprikon {

//------------------------------------------------------------------------
// LitleleprikonDelayController Implementation
//------------------------------------------------------------------------
tresult PLUGIN_API LitleleprikonDelayController::initialize (FUnknown* context)
{
	// Here the Plug-in will be instantiated

	//---do not forget to call parent ------
	tresult result = EditControllerEx1::initialize (context);
	if (result != kResultOk)
	{
		return result;
	}
    Steinberg::Vst::RangeParameter* newParam = new Steinberg::Vst::RangeParameter(STR16("Delay"), kDelayId, STR16("Seconds"), 0.0, 5.0, 1.0, 0.0, Steinberg::Vst::ParameterInfo::kCanAutomate);
            parameters.addParameter(newParam);
//    parameters.addParameter (STR16 ("Bypass"), nullptr, 1, 0, Steinberg::Vst::ParameterInfo::kCanAutomate|Steinberg::Vst::ParameterInfo::kIsBypass, kBypassId);
//    parameters.addParameter (STR16 ("Delay"), STR16 ("sec"), 0, 1, Steinberg::Vst::ParameterInfo::kCanAutomate, kDelayId);


	// Here you could register some parameters

	return result;
}

//------------------------------------------------------------------------
tresult PLUGIN_API LitleleprikonDelayController::terminate ()
{
	// Here the Plug-in will be de-instantiated, last possibility to remove some memory!

	//---do not forget to call parent ------
	return EditControllerEx1::terminate ();
}

//------------------------------------------------------------------------
tresult PLUGIN_API LitleleprikonDelayController::setComponentState (IBStream* state)
{
	// Here you get the state of the component (Processor part)
	if (!state)
		return kResultFalse;

	return kResultOk;
}

//------------------------------------------------------------------------
tresult PLUGIN_API LitleleprikonDelayController::setState (IBStream* state)
{
	// Here you get the state of the controller

	return kResultTrue;
}

//------------------------------------------------------------------------
tresult PLUGIN_API LitleleprikonDelayController::getState (IBStream* state)
{
	// Here you are asked to deliver the state of the controller (if needed)
	// Note: the real state of your plug-in is saved in the processor

	return kResultTrue;
}

//------------------------------------------------------------------------
IPlugView* PLUGIN_API LitleleprikonDelayController::createView (FIDString name)
{
	// Here the Host wants to open your editor (if you have one)
	if (FIDStringsEqual (name, Vst::ViewType::kEditor))
	{
		// create your editor here and return a IPlugView ptr of it
        return nullptr;
	}
	return nullptr;
}

//------------------------------------------------------------------------
} // namespace Litleleprikon
