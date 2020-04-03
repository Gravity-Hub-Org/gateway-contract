const wvs = 10 ** 8;

describe('Deploy script', async function () {
    let seed = ""

    it('Deploy contract', async function () {
        const setScriptTx = setScript({ script: compile(file("../script/gateway.ride")), fee: 1400000,}, seed); 

     /*   const signedIssueTx = issue({
            name: 'SUSDT',
            description: 'USDN Susy token',
            quantity: "10000000000000000000000000",
            decimals: 6,
            fee: 100400000
        }, seed)

        await broadcast(signedIssueTx)*/
        await broadcast(setScriptTx)

        const constructorData = data({
            data: [
                { key: "asset_id", value: "DG2xFkPdDwKUoBkzGAhQtLpSGzfXLiCYPEzeKH2Ad24p" },
                { key: "admins", value: "3q417BA7G3JrRAHAJtWYSTBtQ4XB4dqGMPDwakHWXjy5,BfmQxoMrqXgUHQpVHS7TAy5vPYhmiyD6qMwej34TUz1d" +
                ",AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV"},
                
                { key: "bftCoefficient", value: 2 }
            ],
            fee: 500000
        }, seed);
        await broadcast(constructorData)
    })
})