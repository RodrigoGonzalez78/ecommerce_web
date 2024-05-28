<section class="container  my-5">
  
    <!-- php $validación = \Config\Services::validación(); Esto carga automáticamente el archivo Config\Validation que contiene configuraciones para incluir múltiples conjuntos de reglas -->
    <?php $validation = \Config\Services::validation(); ?>

    <div class="row justify-content-sm-center h-100">
        <div class="col-xxl-4 col-xl-5 col-lg-5 col-md-7 col-sm-12">
            <div class="card shadow-lg">
                <div class="card-body p-5 ">
                    <h1 class="fs-4 card-title fw-bold mb-4 text-center">Actualisar Usuario</h1>

                    <form method="post" action="<?php echo base_url('/update-user?id=' . $user['id']) ?>"  >

                        <div class="mb-2">
                            <label for="name" class="form-label">Nombre</label>
                            <!-- ingreso sel nombre-->
                            <input name="name" type="text" class="form-control" value="<?php echo set_value('name',$user['name']) ?>" placeholder="Nombre">
                            <!-- Error -->
                            <?php if ($validation->getError('name')) { ?>
                                <div class='alert alert-danger mt-2'>
                                    <?= $error = $validation->getError('name'); ?>
                                </div>
                            <?php } ?>
                        </div>
                        <div class="mb-3">
                            <label for="last_name" class="form-label">Apellido</label>
                            <input type="text" name="last_name" class="form-control" value="<?php echo set_value('last_name',$user['last_name']) ?>" placeholder="Apellido">
                            <!-- Error -->
                            <?php if ($validation->getError('last_name')) { ?>
                                <div class='alert alert-danger mt-2'>
                                    <?= $error = $validation->getError('last_name'); ?>
                                </div>
                            <?php } ?>
                        </div>
                        <div class="mb-3">
                            <label for="email" class="form-label">Correo electronico</label>
                            <input name="email" type="email" class="form-control" value="<?php echo set_value('email',$user['email']) ?>" placeholder="correo@algo.com">
                            <!-- Error -->
                            <?php if ($validation->getError('email')) { ?>
                                <div class='alert alert-danger mt-2'>
                                    <?= $error = $validation->getError('email'); ?>
                                </div>
                            <?php } ?>
                        </div>

                        <div class="input-group mb-3">
                            <div class="input-group-prepend">
                                <label class="input-group-text" for="id_profile">Perfil</label>
                            </div>
                            <select class="custom-select" name="id_profile">
                                <option selected>Elija una opcion</option>
                                <?php foreach ($profiles as $profile) { ?>
                                    <option value="<?php echo $profile['id']; ?>" <?php if ($profile['id'] == $user['id_profile']) echo 'selected'; ?>>
                                        <?php echo $profile['descrition']; ?>
                                    </option>
                                <?php } ?>
                            </select>
                            <?php if ($validation->getError('id_profile')) { ?>
                                <div class='alert alert-danger mt-2'>
                                    <?= $error = $validation->getError('id_profile'); ?>
                                </div>
                            <?php } ?>
                        </div>

                       

                        <div class="d-flex align-items-center">
                        <div class="col-1">
                                <a href="<?php echo base_url("/users") ?>"><button type="button" class="btn btn-danger">Cancelar</button></a>
                            </div>
                            <input type="submit" value="Actualisar" class="btn my-btn-primary ms-auto fw-bold">

                        </div>


                    </form>

                </div>
            

            </div>
        </div>

</section>