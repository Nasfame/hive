// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./HiveStorage.sol";

contract HiveStorageTestable is HiveStorage {
  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
