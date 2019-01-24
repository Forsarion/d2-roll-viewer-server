"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var world = require('../bungie/world.json');
var profile = require('../bungie/profile.json');
var response = profile.Response;
var itemId = "6917529084081552355";
var item = response.profileInventory.data.items.find(function (item) {
    return item.itemInstanceId == itemId;
});
var hash = item.itemHash;
var itemsDefinition = world['DestinyInventoryItemDefinition'];
var itemDefinition = itemsDefinition[hash];
var sockets = response.itemComponents.sockets.data[itemId];
console.log(JSON.stringify(sockets, null, 2));
//# sourceMappingURL=index.js.map