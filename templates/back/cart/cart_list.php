<div>

	<?php if (session()->getFlashdata('success')) {
		echo "
      <div class='mt-3 mb-3 ms-3 me-3 h4 text-center alert alert-success alert-dismissible'>
      <button type='button' class='btn-close' data-bs-dismiss='alert'></button>" . session()->getFlashdata('success') . "
  </div>";
	} else if (session()->getFlashdata('error')) {
		echo "
      <div class='mt-3 mb-3 ms-3 me-3 h4 text-center alert alert-danger alert-dismissible'>
      <button type='button' class='btn-close' data-bs-dismiss='alert'></button>" . session()->getFlashdata('error') . "
  </div>";
	} ?>
</div>

<?php
if (!$products) {

?>

    <div class="container py-5">
        <div class="row justify-content-center align-items-center">
            <div class="col-12 col-sm-6 text-center">

                <img src="assets/img/empty_cart.svg" alt="" class="img-fluid">
                <h1 class="mt-5">Carrito vacío</h1>

            </div>
        </div>
    </div>
<?php

} else {


?>
    <section class="h-100">
        <div class="container h-100 py-5">
            <div class="row d-flex justify-content-center align-items-center h-100">
                <div class="col-10">

                    <div class="d-flex justify-content-between align-items-center mb-4">
                        <h3 class="fw-normal mb-0 text-black">Carrito de compras</h3>

                    </div>



                    <?php foreach ($products as $product) : ?>
                        <div class="card rounded-3 mb-3">
                            <div class="card-body p-4">
                                <div class="row d-flex justify-content-between align-items-center">


                                    <div class="col-md-2 col-lg-2 col-xl-2">
                                        <img src="<?php echo base_url('assets/uploads/' . $product['image']); ?>" class="img-fluid rounded-2" alt="">
                                    </div>

                                    <div class="col-md-2 col-lg-3 col-xl-3 ">
                                        <p class="lead fw-normal mb-2"><?php echo $product['name']; ?></p>
                                    </div>

                                    <div class="col-md-2 col-lg-2 col-xl-2 ">

                                        <h5 class="mb-0"><?php echo $product['qty']; ?></h5>
                                    </div>




                                    <div class="col-md-2 col-lg-2 col-xl-2 ">
                                        <h5 class="mb-0"><?php echo '$' . $product['price'] * $product['qty']; ?></h5>
                                    </div>

                                    <div class="col-md-1 col-lg-2 col-xl-2 d-flex">

                                        <a class="px-1" href="<?php echo base_url('/sum-cart-element?rowid=' . $product['rowid'] . '&count=' . $product['qty']) ?>">
                                            <button class="btn my-btn-primary">
                                            +
                                            </button>
                                        </a>
                                        <a href="<?php echo base_url('/rest-cart-element?rowid=' . $product['rowid'] . '&count=' . $product['qty']) ?>">
                                            <button class="btn btn-danger">
                                            -
                                            </button>
                                        </a>

                                        <a class="px-1" href="<?php echo base_url('/remove-cart-element?rowid=' . $product['rowid']) ?>">
                                            <button class="btn btn-danger">
                                                <img src="assets/img/icons/delete.svg" width="20px" height="20px" alt="">
                                            </button>
                                        </a>

                                    </div>

                                </div>
                            </div>
                        </div>




                    <?php endforeach; ?>
                    <div class="card">


                        <div class="card-body d-flex justify-content-between">
                            <a href="<?php echo base_url('/clear-cart') ?>">
                                <button type="button" class="btn btn-danger btn-block btn-lg">Vaciar Carrito</button>
                            </a>

                            <a href="<?php echo base_url('/create-sale') ?>">
                                <button type="button" class="btn my-btn-primary btn-block btn-lg">Terminar Compra</button>
                            </a>
                        </div>
                    </div>

                </div>



            </div>
        </div>
    </section>

<?php
}


?>