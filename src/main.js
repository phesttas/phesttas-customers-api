import express from "express"
import isLoggedIn from './middlewares/login.js'

const app = express()

app.get('/', isLoggedIn, (req, res) => {
  res.json({ "hello": "world" })
})

app.listen(3000, () => {
  console.log("Listen on http://localhost:3000")
})
