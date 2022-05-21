import express from "express"
import { Request, Response } from 'express'
import isLoggedIn from './middlewares/login'

const app = express()

app.get('/', isLoggedIn, (req: Request, res: Response) => {
  res.json({ "hello": "world" })
})

app.listen(3000, () => {
  console.log("Listen on http://localhost:3000")
})
