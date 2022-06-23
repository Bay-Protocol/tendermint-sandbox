const axios = require('axios')

async function main () {
  const args = process.argv.slice(2)
  const from = Number(args[0]) || 0
  const to = Number(args[1]) || 0

  console.log(from, to)
  for (let i = from; i < to; i++) {
    axios.get(`http://localhost:26657/broadcast_tx_commit?tx="${i}"`)
      .then(resp => console.log(resp.data))
      .catch(resp => console.error(resp.response.data))
  }
}

main()
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
