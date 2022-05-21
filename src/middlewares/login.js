const isLoggerIn = (req, res, next) => {
  console.log("middleware login")
  next()
}

export default isLoggerIn
