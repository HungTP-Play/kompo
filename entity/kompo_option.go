package entity

type KompoOption struct {
	In         string // Input , 			default=./docker-compose.yaml | ./docker-compose.yml
	Out        string // Output,			default=helm
	K8sVersion string // K8sVersion, 		default=latest
	Mode       string // Mode, 				default=apply
	Reserve    bool   // Reserve convert 	default=false
}
