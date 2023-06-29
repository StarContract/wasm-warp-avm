const { interactCropImg } = require('../scripts/interact-cropImg');

interactCropImg(
  'localhost',
  1984,
  'http',
  'local',
  'deploy/local/wallet_local.json'
).finally();
