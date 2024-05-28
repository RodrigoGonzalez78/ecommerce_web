<div>
    <!--recuperamos datos con la función Flashdata para mostrarlos-->
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


<!-- php $validación = \Config\Services::validación(); Esto carga automáticamente el archivo Config\Validation que contiene configuraciones para incluir múltiples conjuntos de reglas -->
<?php $validation = \Config\Services::validation(); ?>

<section class="container mt-4 ms-auto">
    <div class="row">
        <div class="card  shadow-lg col-md-10 ">
            <div class="card-body p-3">

                <h1 class=" fs-4 card-title fw-bold mb-4 text-center">Mis Datos</h1>
                <form method="post" action="<?php echo base_url('/update-user-data') ?>">

                    <div class="mb-2">
                        <label for="name" class="form-label">Nombre</label>
                        <!-- ingreso sel nombre-->
                        <input name="name" type="text" class="form-control" value="<?php echo set_value('name', $user['name']) ?>" placeholder="Nombre">
                        <!-- Error -->
                        <?php if ($validation->getError('name')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('name'); ?>
                            </div>
                        <?php } ?>
                    </div>
                    <div class="mb-3">
                        <label for="last_name" class="form-label">Apellido</label>
                        <input type="text" name="last_name" class="form-control" value="<?php echo set_value('last_name', $user['last_name']) ?>" placeholder="Apellido">
                        <!-- Error -->
                        <?php if ($validation->getError('last_name')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('last_name'); ?>
                            </div>
                        <?php } ?>
                    </div>
                    <div class="mb-3">
                        <label for="email" class="form-label">Correo electronico</label>
                        <input name="email" type="email" class="form-control" value="<?php echo set_value('email', $user['email']) ?>" placeholder="correo@algo.com">
                        <!-- Error -->
                        <?php if ($validation->getError('email')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('email'); ?>
                            </div>
                        <?php } ?>
                    </div>


                    <div class="d-flex align-items-center">

                        <input type="submit" value="Actualisar" class="btn my-btn-primary ms-auto fw-bold">

                    </div>


                </form>

                <h1 class=" fs-4 card-title fw-bold mb-4 text-center">Cambiar Contraseña</h1>
                <form method="post" action="<?php echo base_url('/update-password') ?>">


                    <div class="mb-3">
                        <label for="password" class="form-label">Contraseña</label>
                        <input name="password" type="password" class="form-control" placeholder="Contraseña (mínimo 8 caracteres)">
                        <!-- Error -->
                        <?php if ($validation->getError('password')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('password'); ?>
                            </div>
                        <?php } ?>
                    </div>

                    <div class="d-flex align-items-center">

                        <input type="submit" value="Cambiar" class="btn my-btn-primary ms-auto fw-bold">

                    </div>


                </form>
                <h1 class=" fs-4 card-title fw-bold mb-4 text-center">Mi dirección</h1>
                <form method="post" action="<?php echo base_url('/update-address') ?>">


                    <div class="mb-3">
                        <label for="street" class="form-label">Calle</label>
                        <input name="street" type="text" class="form-control" value="<?php echo set_value('street', $address['street']) ?>" placeholder="Calle">
                        <!-- Error -->
                        <?php if ($validation->getError('street')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('street'); ?>
                            </div>
                        <?php } ?>
                    </div>
                    <div class="mb-3">
                        <label for="postal_code" class="form-label">Código Postal</label>
                        <input name="postal_code" type="text" class="form-control" value="<?php echo set_value('postal_code', $address['postal_code']) ?>" placeholder="Código Postal">
                        <!-- Error -->
                        <?php if ($validation->getError('postal_code')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('postal_code'); ?>
                            </div>
                        <?php } ?>
                    </div>
                    <div class="mb-3">
                        <label for="neighborhood" class="form-label">Barrio</label>
                        <input name="neighborhood" type="text" class="form-control" value="<?php echo set_value('neighborhood', $address['neighborhood']) ?>" placeholder="Barrio">
                        <!-- Error -->
                        <?php if ($validation->getError('neighborhood')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('neighborhood'); ?>
                            </div>
                        <?php } ?>
                    </div>
                    <div class="mb-3">
                        <label for="city" class="form-label">Ciudad</label>
                        <input name="city" type="text" class="form-control" value="<?php echo set_value('city', $address['city']) ?>" placeholder="Ciudad">
                        <!-- Error -->
                        <?php if ($validation->getError('city')) { ?>
                            <div class='alert alert-danger mt-2'>
                                <?= $error = $validation->getError('city'); ?>
                            </div>
                        <?php } ?>
                    </div>

                    <div class="d-flex align-items-center">

                        <input type="submit" value="Actualisar" class="btn my-btn-primary ms-auto fw-bold">

                    </div>


                </form>



            </div>

        </div>

    </div>
</section>