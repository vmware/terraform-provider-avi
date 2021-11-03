---
- hosts: localhost
  connection: local
  gather_facts: no
  roles:
    - role: avinetworks.avisdk
  vars:
    controller: "{{ ansible_host }}"
    username: admin
    cloud_name: "Default-Cloud"
    ansible_become: yes
    ansible_become_password: "{{ password }}"
    avi_version: ${avi_version}
    tenant_name: "admin"
    
  tasks:
    - name: Remove all DNS Service Refs from System Configuration
      avi_api_session:
        controller: "{{ controller }}"
        username: "{{ username }}"
        password: "{{ password }}"
        http_method: patch
        path: "systemconfiguration"
        tenant: "admin"
        data:
          replace:
            dns_virtualservice_refs: ""

    - name: Get Virtual Service Information
      avi_api_session:
        controller: "{{ controller }}"
        username: "{{ username }}"
        password: "{{ password }}"
        http_method: get
        path: virtualservice
        tenant: "*"
        params:
          fields: "name,enabled,uuid,tenant_ref"
      register: vs_results

    - name: Display all Virtual Services
      ansible.builtin.debug:
        var: vs_results.obj.results

    - name: Delete all Virtual Services
      avi_api_session:
        controller: "{{ controller }}"
        username: "{{ username }}"
        password: "{{ password }}"
        http_method: delete
        path: "virtualservice/{{ item.uuid }}"
        tenant: "*"
      loop: "{{ vs_results.obj.results }}"

    - name: Get Service Engine Information
      avi_api_session:
        controller: "{{ controller }}"
        username: "{{ username }}"
        password: "{{ password }}"
        http_method: get
        path: serviceengine
        tenant: "*"
        params:
          fields: "name,enabled,uuid,cloud_ref"
      register: se_results

    - name: Display all Service Engines
      ansible.builtin.debug:
        var: se_results.obj.results

    - name: Delete all Service Engines
      avi_api_session:
        controller: "{{ controller }}"
        username: "{{ username }}"
        password: "{{ password }}"
        http_method: delete
        path: "serviceengine/{{ item.uuid }}"
      loop: "{{ se_results.obj.results }}"