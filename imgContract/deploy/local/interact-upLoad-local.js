const { interactUpLoad } = require('../scripts/interact-upLoad');

interactUpLoad(
  'localhost',
  1984,
  'http',
  'local',
  'deploy/local/wallet_local.json'
).finally();
