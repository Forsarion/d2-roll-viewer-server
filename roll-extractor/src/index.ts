import { Response } from "./Response";

const world = require('../bungie/world.json')
const profile = require('../bungie/profile.json')

const response: Response = profile.Response
const itemId = "6917529084081552355"

const item = response.profileInventory.data.items.find(item => {
    return item.itemInstanceId == itemId
})!

const hash = item.itemHash
const itemsDefinition = world['DestinyInventoryItemDefinition']
const itemDefinition = itemsDefinition[hash]
const sockets = response.itemComponents.sockets.data[itemId]!

console.log(JSON.stringify(sockets, null, 2))