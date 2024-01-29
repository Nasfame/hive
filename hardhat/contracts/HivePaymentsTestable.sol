// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./HivePayments.sol";

contract HivePaymentsTestable is HivePayments {
  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
