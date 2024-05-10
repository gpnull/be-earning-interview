const AttendanceTracker = artifacts.require("AttendanceTracker");

module.exports = function(deployer) {
    deployer.deploy(AttendanceTracker);
};