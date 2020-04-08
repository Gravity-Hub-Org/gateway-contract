const Supersymmetry = artifacts.require("./Supersymmetry.sol");

module.exports = async function(deployer, network, accounts) {
    let ethAssetId = "89LVNfmKavK1zNpKswUo3rJ72GJfyJcYPij6DMvxfr4S"
    await deployer.deploy(Supersymmetry, 
        ["0x25513f13e9a69dcf2925BDfcED50C13709554358", "0xe2AD6550287D3AB7AA0cf44A0a15B1946C0D8De5", "0x25513f13e9a69dcf2925BDfcED50C13709554358", "0x25513f13e9a69dcf2925BDfcED50C13709554358", "0x25513f13e9a69dcf2925BDfcED50C13709554358"], 
        2, ethAssetId, 8);
};