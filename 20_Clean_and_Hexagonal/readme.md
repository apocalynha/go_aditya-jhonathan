# Summary

1. Good Architecture adalah prinsip pemisahan perhatian (separation of concern) dengan menggunakan lapisan (layers) untuk membangun aplikasi yang modular, dapat diperluas, mudah dirawat, dan dapat diuji. Ini berarti memecah aplikasi menjadi komponen-komponen terpisah dengan tugas dan tanggung jawab yang jelas, dan menjaga agar setiap komponen tersebut fokus pada hal-hal yang relevan.

2. Clean Architecture

    Clean Architecture adalah konsep desain yang diperkenalkan oleh Robert C. Martin. Ia berfokus pada pemisahan konsep-konsep inti dalam aplikasi dan mengaturnya dalam lapisan yang dapat diuji, dipecah, dan diubah tanpa mengganggu bagian lain dari aplikasi. Ada empat lapisan utama dalam Clean Architecture:

    - Entities: Ini adalah entitas dasar yang mewakili objek bisnis dalam aplikasi.
    - Use Cases (Interactors): Ini adalah aturan bisnis dan logika aplikasi yang berhubungan dengan penggunaan entitas.
    - Interface Adapters: Ini berfungsi untuk menghubungkan Use Cases dengan input dan output. Termasuk hal seperti pengontrol, presenter, dan gateways ke database atau sumber data eksternal.
    - Frameworks & Drivers: Ini adalah bagian yang berhubungan dengan kerangka kerja dan sumber daya eksternal seperti database, web, UI, dll.

    Clean Architecture mengutamakan prinsip "Dependency Inversion," yang mengarah pada pemisahan kode tingkat tinggi dari kode tingkat rendah dan pembalikannya, sehingga kode tingkat tinggi tidak tergantung pada kode tingkat rendah. Ini memungkinkan aplikasi untuk lebih mudah diuji dan dikembangkan ulang.

3. Hexagonal Architecture

    Hexagonal Architecture, juga dikenal sebagai Ports and Adapters, adalah pendekatan desain yang menekankan pada pemisahan komponen aplikasi ke dalam dua konsep utama: Ports dan Adapters. Ini membagi aplikasi menjadi dua bagian utama:

    - Port: Ini adalah antarmuka yang mendefinisikan bagaimana aplikasi berinteraksi dengan dunia luar. Port ini tidak berisi logika bisnis, tetapi hanya mendefinisikan apa yang dapat dilakukan oleh aplikasi.
    - Adapter: Ini adalah implementasi konkret dari Port yang menghubungkan aplikasi dengan sumber data eksternal atau komponen lainnya. Adapter ini mengimplementasikan logika bisnis dan menyediakan akses ke sumber daya eksternal.

    Hexagonal Architecture mengutamakan prinsip "Dependency Injection," yang mengizinkan komponen aplikasi untuk bergantung pada antarmuka (Port) dan bukan implementasi konkretnya (Adapter). Ini membuat aplikasi lebih fleksibel, mudah diuji, dan dapat diganti.