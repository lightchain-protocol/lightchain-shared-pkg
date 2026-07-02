// Package bindings contains Go bindings generated from Solidity contract ABIs
// via abigen. These bindings enable type-safe interaction with the LightChain
// protocol contracts from Go code.
//
// The ABI JSON files in pkg/chain/abis/ are extracted from the concrete Foundry
// build artifacts in contracts/out/. Concrete contract ABIs are required here
// because the interface artifacts omit inherited Ownable/UUPS methods such as
// owner(), initialize(), and upgradeToAndCall(). Custom Solidity error entries
// must still be stripped because abigen v1.10.x does not support them.
//
// Regenerate with:
//
//	go generate ./pkg/chain/bindings/...
package bindings

//go:generate sh -c "t=$$(mktemp) && trap 'rm -f \"$$t\"' EXIT && python3 -c \"import json,sys; print(json.dumps([e for e in json.load(sys.stdin) if e.get('type')!='error']))\" < ../abis/IWorkerRegistry.json > \"$$t\" && abigen --abi=\"$$t\" --pkg=bindings --type=WorkerRegistry --out=worker_registry.go"
//go:generate sh -c "t=$$(mktemp) && trap 'rm -f \"$$t\"' EXIT && python3 -c \"import json,sys; print(json.dumps([e for e in json.load(sys.stdin) if e.get('type')!='error']))\" < ../abis/IJobRegistry.json > \"$$t\" && abigen --abi=\"$$t\" --pkg=bindings --type=JobRegistry --out=job_registry.go"
//go:generate sh -c "t=$$(mktemp) && trap 'rm -f \"$$t\"' EXIT && python3 -c \"import json,sys; print(json.dumps([e for e in json.load(sys.stdin) if e.get('type')!='error']))\" < ../abis/IAIConfig.json > \"$$t\" && abigen --abi=\"$$t\" --pkg=bindings --type=AIConfig --out=ai_config.go"
//go:generate sh -c "t=$$(mktemp) && trap 'rm -f \"$$t\"' EXIT && python3 -c \"import json,sys; print(json.dumps([e for e in json.load(sys.stdin) if e.get('type')!='error']))\" < ../abis/IReputationRegistry.json > \"$$t\" && abigen --abi=\"$$t\" --pkg=bindings --type=ReputationRegistry --out=reputation_registry.go"
