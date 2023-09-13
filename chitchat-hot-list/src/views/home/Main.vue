<template>
      <main>
        <div class="mx-auto max-w-full py-6 sm:px-6 lg:px-8">
          <!-- Your content -->
            
            <!-- <div class="bg-white"> -->
            <div>
                <!-- <div class="mx-auto max-w-full px-4 py-16 sm:px-6 sm:py-24 lg:max-w-full lg:px-8" > -->
                <div class="mx-auto max-w-full lg:max-w-full" >
                    <h2 class="sr-only">产品</h2>
            
                    <div class="grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 xl:gap-x-8">
                       
                            <div v-for="product in products" class="group bg-gray-200 " >
                                <div  class="aspect-h-1 aspect-w-1 w-full  overflow-hidden rounded-lg  xl:aspect-h-8 xl:aspect-w-7">
                                 
                                  <el-scrollbar height="400px" class="py-2">
                                      <div v-for="product in products" :key="product.id" :href="product.href"  class="h-full w-full px-5 py-1 cursor-pointer " > 
                                         <a class="mt-4 text-sm font-medium tracking-wide text-blue-700 hover:text-blue-400 ">{{product.id}}.{{ product.name }}</a>
                                      </div>
                                    </el-scrollbar>
                                  <!-- <img :src="product.imageSrc" :alt="product.imageAlt" class="h-full w-full object-cover object-center group-hover:opacity-75" /> -->
                                </div>
                                <!-- <h3 class="mt-4 text-sm text-gray-700">{{ product.name }}</h3> -->
                                <!-- <p class="mt-1 text-lg font-medium text-gray-900">{{ product.price }}</p> -->
                            
                           
                              </div>
                      
                            
                         
                    </div>
                </div>
            </div>  

        </div>
      </main>
</template>


<!--
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/aspect-ratio'),
    ],
  }
  ```
-->

  

  
  <script setup>
 
  const products = [
    {
      id: 1,
      name: 'Earthen Bottle 数据在各领域都有着十分重要的作用。',
      href: '#',
      price: '$48',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-01.jpg',
      imageAlt: 'Tall slender porcelain bottle with natural clay textured body and cork stopper.',
    },
    {
      id: 2,
      name: 'Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据需要灵活选择',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-02.jpg',
      imageAlt: 'Olive drab green insulated bottle with flared screw lid and flat top.',
    },
    {
      id: 3,
      name: 'Focus Paper Refill',
      href: '#',
      price: '$89',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-03.jpg',
      imageAlt: 'Person using a pen to cross a task off a productivity paper card.',
    },
    {
      id: 4,
      name: 'Machined Mechanical PenciNomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据需要灵活选择Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据l',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-04.jpg',
      imageAlt: 'Hand holding black machined steel mechanical pencil with brass tip and top.',
    },
    {
      id: 4,
      name: 'Machined Mechanical Pencil',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-04.jpg',
      imageAlt: 'Hand holding black machined steel mechanical pencil with brass tip and top.',
    },
    {
      id: 4,
      name: 'Machined Mechanical Pencil',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-04.jpg',
      imageAlt: 'Hand holding black machined steel mechanical pencil with brass tip and top.',
    },
     {
      id: 4,
      name: 'Machined Mechanical PenciNomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据l',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-04.jpg',
      imageAlt: 'Hand holding black machined steel mechanical pencil with brass tip and top.',
    },
    {
      id: 2,
      name: 'Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据需要灵活选择',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-02.jpg',
      imageAlt: 'Olive drab green insulated bottle with flared screw lid and flat top.',
    },
    {
      id: 2,
      name: 'Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据需要灵活选择Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据Nomad Tumbler数据在各领域都有着十分重要的作用。数据获取的方法有很多，可以根据aaaaa',
      href: '#',
      price: '$35',
      imageSrc: 'https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-02.jpg',
      imageAlt: 'Olive drab green insulated bottle with flared screw lid and flat top.',
    },
    
    // More products...
  ]
  </script>