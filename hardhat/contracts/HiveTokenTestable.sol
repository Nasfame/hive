// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./HiveToken.sol";

// a version of HiveToken.sol that can be called by any address
// so we can run unit tests
contract HiveTokenTestable is HiveToken {
  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) HiveToken(name, symbol, initialSupply) {}

  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
