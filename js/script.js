// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file generates user input
function enterClicked() {
  // input
  const streetName = document.getElementById("streetName").value
  const streetNumber = parseInt(document.getElementById("streetNumber").value)

  // output
  document.getElementById("address").innerHTML =
    "Your address is: " + streetNumber + " " + streetName + "."
}
