const wvs = 10 ** 8;

describe('Deploy script', async function () {
    before(async function () {
        await setupAccounts({
            contract: 100000 * wvs
        })
    });

    it('Issue contract', async function () {
        const signedIssueTx = issue({
            name: 'sEth',
            description: 'eth Susy token',
            quantity: "11043650500000000",
            decimals: 8,
            fee: 100400000
        }, accounts.contract)

        await broadcast(signedIssueTx)
        console.log(signedIssueTx.id)
        console.log(accounts.contract)
    })
})