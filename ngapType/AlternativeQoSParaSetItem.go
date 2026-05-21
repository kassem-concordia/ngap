package ngapType


// One alternative GBR QoS parameter set within AlternativeQoSParaSetList. //kassem
type AlternativeQoSParaSetItem struct { //kassem
	AlternativeQoSParaSetIndex AlternativeQoSParaSetIndex                              
	GuaranteedFlowBitRateDL    *BitRate                                                `aper:"optional"` 
	GuaranteedFlowBitRateUL    *BitRate                                                `aper:"optional"` 
	MaximumFlowBitRateDL       *BitRate                                                `aper:"optional"` 
	MaximumFlowBitRateUL       *BitRate                                                `aper:"optional"` 
	PacketDelayBudget          *PacketDelayBudget                                      `aper:"optional"` 
	PacketErrorRate            *PacketErrorRate                                        `aper:"optional"` 
	IEExtensions               *ProtocolExtensionContainerAlternativeQoSParaSetItemExtIEs `aper:"optional"` 
} //kassem