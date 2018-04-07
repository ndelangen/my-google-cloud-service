const Datastore = require('@google-cloud/datastore')

module.exports = projectId => {
  const datastore = new Datastore({ projectId })

  return {
    saveOrganisation: organisation => datastore
      .save({
        key: datastore.key(['organsitions', 'uniqueId-' + Date.now()]),
        data: { github_organisation: organisation },
      })
  }
}

