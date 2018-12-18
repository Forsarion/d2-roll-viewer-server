export interface Response {
    profileInventory: {
        data: {
            items: [
                {
                    itemHash: number
                    itemInstanceId: string           
                }
            ]
        }
    }
    itemComponents: {
        sockets: {
            data: {
                [key: string]: {
                    sockets: [
                        {
                            plugHash: number,
                            isEnabled: boolean,
                            isVisible: boolean,
                            reusablePlugHashes: [
                                number
                            ],
                            reusablePlugs: [
                                {
                                    plugItemHash: number,
                                    canInsert: boolean,
                                    enabled: boolean
                                }
                            ]
                        }
                    ]
                }
            }
        }
    }
}