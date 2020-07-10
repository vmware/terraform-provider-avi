pools = {
  "pool1" = {
    lb_algorithm = "LB_ALGORITHM_ROUND_ROBIN"
    lb_algorithm_consistent_hash_hdr = "hdr1"
    serversList = [
    ]
  }
  "pool2" = {
    lb_algorithm = "LB_ALGORITHM_ROUND_ROBIN"
    lb_algorithm_consistent_hash_hdr = "hdr2"
    serversList = [
      {
        ip = "10.20.20.44"
        port = 80 
      },
      {
        ip = "10.20.20.45"
        port = 81 
      },
      {
        ip = "10.20.20.46"
        port = 82 
      }
    ]
  }
}
