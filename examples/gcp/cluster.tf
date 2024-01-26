resource "kops_cluster" "cluster" {
  name               = local.clusterName
  kubernetes_version = "1.28.5"
  channel            = "stable"

  api {
    access = [
      "0.0.0.0/0",
    ]
    load_balancer {
      type = "Public"
    }
  }

  config_store {
    base = "${var.kops_state_store_bucket}/${local.clusterName}"
  }

  cloud_provider {
    gce {
      project   = var.google_project_id
      multizone = true
      pd_csi_driver {
        enabled = true
      }
    }
  }

  iam {
    legacy                   = false
    allow_container_registry = true
  }

  networking {
    network_id = local.network

    pod_cidr                 = "10.4.0.0/14"
    service_cluster_ip_range = "10.0.48.0/20"

    # enable Google networking mode
    gcp {}

    # subnets
    subnet {
      name   = "kops-0"
      type   = "Public"
      region = var.region
      zone   = "" # Google subnets are regional, not zonal.
    }

    topology {
      dns = "None"
    }
  }

  kube_proxy {
    enabled = true
  }

  # etcd clusters
  etcd_cluster {
    name = "main"
    member {
      name           = "a"
      instance_group = "control-plane-0"
    }
    member {
      name           = "b"
      instance_group = "control-plane-1"
    }
    member {
      name           = "c"
      instance_group = "control-plane-2"
    }
  }

  etcd_cluster {
    name = "events"
    member {
      name           = "a"
      instance_group = "control-plane-0"
    }
    member {
      name           = "b"
      instance_group = "control-plane-1"
    }
    member {
      name           = "c"
      instance_group = "control-plane-2"
    }
  }

  ssh_access = ["0.0.0.0/0"]
}

resource "kops_instance_group" "control-plane-0" {
  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-0"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.controlPlaneType
  subnets      = ["kops-0"]
  zones        = ["us-central1-a"]
}

resource "kops_instance_group" "control-plane-1" {
  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-1"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.controlPlaneType
  subnets      = ["kops-0"]
  zones        = ["us-central1-b"]
}

resource "kops_instance_group" "control-plane-2" {
  cluster_name = kops_cluster.cluster.id
  name         = "control-plane-2"
  role         = "ControlPlane"
  min_size     = 1
  max_size     = 1
  machine_type = local.controlPlaneType
  subnets      = ["kops-0"]
  zones        = ["us-central1-c"]
}

resource "kops_instance_group" "nodes" {
  cluster_name = kops_cluster.cluster.id
  name         = "kops-node"
  role         = "Node"
  min_size     = 3
  max_size     = 6
  machine_type = local.nodeType
  subnets      = ["kops-0"]
  zones        = local.zones
}

resource "kops_cluster_updater" "updater" {
  cluster_name = kops_cluster.cluster.id

  keepers = {
    cluster         = kops_cluster.cluster.revision
    control-plane-0 = kops_instance_group.control-plane-0.revision
    control-plane-1 = kops_instance_group.control-plane-1.revision
    control-plane-2 = kops_instance_group.control-plane-2.revision
    nodes           = kops_instance_group.nodes.revision
  }

  rolling_update {
    skip                = false
    fail_on_drain_error = true
    fail_on_validate    = true
    validate_count      = 1
  }

  validate {
    skip = false
  }
}

data "kops_kube_config" "kube_config" {
  cluster_name = kops_cluster.cluster.id

  depends_on = [kops_cluster_updater.updater]
}
