const express = require('express')
const bodyParser = require('body-parser');

const googleDataStoreClient = require('./googleDataStoreClient')('norbert-de-langen-speeltuin')

const app = express()
app.use(bodyParser.json())

app.post('/api/register', (req, res) => {
  const { organisation } = req.body
  if (organisation) {
    googleDataStoreClient.saveOrganisation(organisation)
      .then(() => res.json(organisation))
      .catch(() => res.status(500).send())
  } else {
    res.status(400).send()
  }
})

app.get('/*', (req, res) => res.status(418).send('I\'m a teapot!'))

if (module === require.main) {
  const server = app.listen(process.env.PORT || 8081, () => {
    const port = server.address().port
    console.log(`App listening on port ${port}`)
  })
}

module.exports = app