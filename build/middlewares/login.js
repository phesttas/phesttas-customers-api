"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const isLoggerIn = (req, res, next) => {
    console.log("middleware login");
    next();
};
exports.default = isLoggerIn;
