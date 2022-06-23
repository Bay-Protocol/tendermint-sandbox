clean:
	rm tendermint-sandbox
	rm -rf ./data
	rm -rf ./tmhome/data/*.db ./tmhome/data/priv_validator_state.json
	cp ./tmhome/data/template_priv_validator_state.json ./tmhome/data/priv_validator_state.json

start:
	go build && ./tendermint-sandbox -config "./tmhome/config/config.toml"

test:
	node ./js/app.js $(from) $(to)
