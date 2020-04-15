const wvs = 10 ** 8;

describe('Deploy script', async function () {
    let seed = "contract#3dc511c"

    it('Deploy contract', async function () {
        const setScriptTx = setScript({ script: compile(file("../script/gateway.ride")), fee: 1400000,}, seed); 
        await broadcast(setScriptTx)

        const constructorData = data({
            data: [
                { key: "admins", value: "AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV" +
                ",AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV,AXbaBkJNocyrVpwqTzD4TpUY8fQ6eeRto9k1m2bNCzXV"},

                { key: "bftCoefficient", value: 1 },
                { key: "rq_timeout", value: 10 }
            ],
            fee: 500000
        }, seed);
        await broadcast(constructorData)
    })
})