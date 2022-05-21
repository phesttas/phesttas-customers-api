"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const login_js_1 = __importDefault(require("./middlewares/login.js"));
const app = (0, express_1.default)();
app.get('/', login_js_1.default, (req, res) => {
    res.json({ "hello": "world" });
});
app.listen(3000, () => {
    console.log("Listen on http://localhost:3000");
});
