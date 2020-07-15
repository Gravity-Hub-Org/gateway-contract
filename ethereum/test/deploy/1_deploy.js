
const Supersymmetry = artifacts.require("Supersymmetry");
let ethAssetId = "DcegJK3HbAKtrqwjQq6AcbGxDNZ8y4obNLJ5SuwiV6dT"

contract("Mint test", async accounts => {
    let susyContract = null;

    before(async () => {
        susyContract = await Supersymmetry.deployed( ["0xe2d3f4c3170abad93339c1251b530b7f87ced3ad", "0xe2AD6550287D3AB7AA0cf44A0a15B1946C0D8De5", 
        "0x25513f13e9a69dcf2925BDfcED50C13709554358", "0x25513f13e9a69dcf2925BDfcED50C13709554358", "0x25513f13e9a69dcf2925BDfcED50C13709554358"], 
        1, ethAssetId, 8);
        console.log("Contract address:" + susyContract.address)
    });

    it("New", async () => {
        let res = await susyContract.newWavesToken({from: accounts[0] })
        console.log("Waves ERC20:" +res.logs[0].args[0])
    });
});