backend "consul" {
    # https://www.vaultproject.io/docs/configuration/storage/consul.html#service_address
  	address = "consul-server-bootstrap:8500"
    service_address = "consul-server-2:8500"
  	advertise_addr = "http://consul-server-bootstrap:8300"
  	scheme = "http"
    consistency_mode = "strong"
}
listener "tcp" {
    address = "0.0.0.0:8200"
    #tls_cert_file = "/config/server.crt"
    #tls_key_file = "/config/server.key"
    tls_disable = 1
}
disable_mlock = true
ui = true
