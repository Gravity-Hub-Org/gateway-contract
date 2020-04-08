const wvs = 10 ** 8;

describe('Deploy script', async function () {
    let wavesERC20Address = ""
    let ethAssetId = "89LVNfmKavK1zNpKswUo3rJ72GJfyJcYPij6DMvxfr4S"
    let seed = "contract#687fd090"

    it('Deploy contract', async function () {
      
        const constructorData = data({
            data: [
                { key: "erc20_address_WAVES", value: wavesERC20Address },
                { key: "asset_type__WAVES", value: 0 },
                { key: "asset_status_WAVES", value: 3 },
                { key: "asset_decimals__WAVES", value: 8 },

                { key: "bftCoefficient", value: 1 },
                { key: "rq_timeout", value: 10 },


                { key: "erc20_address_" + ethAssetId, value: "0x0000000000000000000000000000000000000000" },
                { key: "asset_type_" + ethAssetId, value: 1 },
                { key: "asset_status_" + ethAssetI, value: 3 },
                { key: "asset_decimals__" + ethAssetId, value: 8 },

                { key: "bftCoefficient", value: 1 },
                { key: "rq_timeout", value: 10 }
            ],
            fee: 500000
        }, seed);
        await broadcast(constructorData)
    })
})