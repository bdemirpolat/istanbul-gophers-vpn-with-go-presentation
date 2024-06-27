# kommnunity.com Istanbul Gophers, Writing VPN with Golang Presentation

- Client ve server uygulamaları sadece Linux dağıtımlarında çalışmaktadır, MacOS ve Windows için çalıştırmak isterseniz "sudo ip addr add %s/24 dev %s" ve "sudo ip link set dev %s up" gibi komutları işletim sisteminize göre uygulamalısınız.

- VPN istemcisi tarafındaki YOUR_PUBLIC_VPN_SERVER_ADDR değerini VPN sunusu belirlediğiniz cihazın IP adresi ile değiştirmelisiniz, bu sayede VPN istemcisi VPN sunucusuna TCP bağlantısı yaratabilecektir.

- Server uygulamasını çalıştırdıktan sonra "ifconfig" komutu ile ağ arayüzlerini listeleyebilir ve 10.10.10.2 ip adresine sahip TUN interface'i görebilirsiniz.
- TUN interface başarılı şekilde yaratıldı ise VPN sunucusu içerisinde "httpserver" klasörü altındaki uygulamayı çalıştırabilirsiniz, bu uygulama size 10.10.10.1:8080/test adresinden erişimi test etmeniz için bir HTTP server sunacaktır.
- "client" klasörü altındaki uygulama client sunucusunda çalıştırıldıktan sonra daha önce client üzerinden erişime sahip olmadığınız 10.10.10.2:8080/test adresine VPN sunucusu üzerinden erişebildiğinizi görebileceksiniz.
- Sunucu tarafına debug print'ler ekleyerek VPN sunucusu gelen giden trafiği izleyebilirsiniz.