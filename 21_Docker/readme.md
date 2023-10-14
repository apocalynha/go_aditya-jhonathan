# Summary

1. Docker adalah platform perangkat lunak yang digunakan untuk mengembangkan, menguji, dan menjalankan aplikasi dalam lingkungan yang terkisolasi yang disebut "Container". Docker memungkinkan untuk membuat, mengemas, dan mendistribusikan aplikasi dengan cara yang dapat diandalkan dan mudah dipindahkan antara lingkungan pengembangan dan produksi.

2. Container adalah wadah yang dapat berisi aplikasi berserta semua dependensinya, termasuk perangkat lunak, perpustakaan, dan konfigurasi yang diperlukan untuk menjalankannya. Container bukan virtual mesin karena container merupakan proses dengan file sistem terisolasi yang dibutuhkan aplikasi.

3. Image adalah blueprint aplikasi yang berisi perangkat lunak, perpustakaan, dan konfigurasi yang diperlukan untuk menjalankan aplikasi dalam kontainer. 

4. Dockerfile adalah file teks yang digunakan untuk mendefinisikan langkah-langkah dalam proses pembuatan gambar Docker. Ini adalah alat utama dalam menciptakan gambar yang dapat digunakan untuk membuat kontainer.

5. Hal yang dilakukan pada docker yaitu membuat dockerfile untuk mendeskripsikan langkah membuat image. Kemudian membuat docker image dan container berdasarkan image. Selanjutnya meng-upload atau publish image ke Docker Hub (Image Repository). Terakhir download dan running dari Docker Hub.


# Tugas

1. Install docker & docker compose

![docker](/21_Docker/screenshots/docker.JPG)

2. Create dockerfile.

![dockerfile](/21_Docker/screenshots/dockerfile.JPG)

3. Clone your code and integrate to your dockerfile.

![clonecode](/21_Docker/screenshots/clonecode.JPG)

4. Build your container

![build](/21_Docker/screenshots/build.JPG)

5. Push the image to docker registry

![docker_push](/21_Docker/screenshots/docker_push.JPG)

6. Deploy in your local machine

![deploy](/21_Docker/screenshots/deploy.JPG)

8. Try to Deploy in the Cloud (optional)

![deploy](/21_Docker/screenshots/deploy.JPG)