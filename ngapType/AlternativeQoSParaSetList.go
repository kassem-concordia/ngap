package ngapType


// Ordered list (max 8 entries) of alternative GBR QoS parameter sets. 
type AlternativeQoSParaSetList struct { //kassem
	List []AlternativeQoSParaSetItem `aper:"sizeLB:1,sizeUB:8"` 
} //kassem