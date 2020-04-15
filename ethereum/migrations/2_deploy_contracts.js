const Supersymmetry = artifacts.require("./Supersymmetry.sol");

module.exports = async function(deployer, network, accounts) {
    let ethAssetId = "DcegJK3HbAKtrqwjQq6AcbGxDNZ8y4obNLJ5SuwiV6dT"
    await deployer.deploy(Supersymmetry, 
        ["0xe2d3f4c3170abad93339c1251b530b7f87ced3ad", "0xe2AD6550287D3AB7AA0cf44A0a15B1946C0D8De5", "0x25513f13e9a69dcf2925BDfcED50C13709554358", "0x25513f13e9a69dcf2925BDfcED50C13709554358", "0x25513f13e9a69dcf2925BDfcED50C13709554358"], 
        1, ethAssetId, 8);
};