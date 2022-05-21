import { Request, Response, NextFunction } from 'express'

const isLoggerIn = (req: Request, res: Response, next: NextFunction) => {
  console.log("middleware login")
  next()
}

export default isLoggerIn
