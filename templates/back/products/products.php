<div>
    <?php if (session()->getFlashdata('success')) {
        echo "
      <div class='mt-3 mb-3 ms-3 me-3 h4 text-center alert alert-success alert-dismissible'>
      <button type='button' class='btn-close' data-bs-dismiss='alert'></button>" . session()->getFlashdata('success') . "
  </div>";
    }
    ?>
</div>

<?php $validation = \Config\Services::validation(); ?>


<div class="row m-2">


    <div class="col-3">

    
        <form  method="post" action="<?php echo base_url('/filter-products') ?>" enctype="multipart/form-data">


            <div class="col mb-2">
                <label for="search" class="form-label">Buscar</label>
                <!-- ingreso sel nombre-->
                <input name="search" type="text" class="form-control" value="<?php echo set_value('search') ?>" placeholder="Palabra">
                <!-- Error -->
                <?php if ($validation->getError('search')) { ?>
                    <div class='alert alert-danger mt-2'>
                        <?= $error = $validation->getError('search'); ?>
                    </div>
                <?php } ?>
            </div>

            <div class="col input-group mb-3">
                <div class="input-group-prepend">
                    <label class="input-group-text" for="id_categorie">Categoria</label>
                </div>
                <select class="custom-select" name="id_categorie">
                    <option selected>Todos</option>
                    <?php foreach ($categories as $category) { ?>
                        <option value="<?php echo $category['id']; ?>"><?php echo $category['name']; ?></option>
                    <?php } ?>
                </select>
                <?php if ($validation->getError('id_categorie')) { ?>
                    <div class='alert alert-danger mt-2'>
                        <?= $error = $validation->getError('id_categorie'); ?>
                    </div>
                <?php } ?>
            </div>
            <div class=" col">
                <input type="submit" value="Filtrar" class="btn fw-bold my-btn-primary ms-auto">
            </div>


        </form>
    </div>
    <?php
    if (session()->get('id_profile') == 1) {
    ?>


        <div class="col-auto ms-auto">
            <a href="<?php echo base_url("/new-product") ?>" class="text-dark"><button type="button" class="btn my-btn-primary fw-bold">Agregar</button></a>

        </div>
        <div class="col-auto">
            <a href="<?php echo base_url("/disableds-products") ?>"><button type="button" class="btn my-btn-primary fw-bold">Inactivos</button></a>

        </div>



    <?php

    } ?>



</div>

<hr>

<div class="row m-2">


    <?php

    foreach ($products as $product) {
    ?>
        <!-- Incluir la vista product_card para cada producto -->
        <?php echo view('back/products/component/product_card', ['product' => $product]); ?>
    <?php
    }

    ?>



</div>
</div>