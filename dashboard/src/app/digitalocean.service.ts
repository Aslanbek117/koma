import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { Observable } from "rxjs/Rx";
import { StoreService } from './store.service';
import { localAddress } from '../environments/environment'

@Injectable()
export class DigitaloceanService {
  private BASE_URL = '/digitalocean'

  constructor(private http: Http, private storeService: StoreService) { }

  public getProfile(){
    return this.http
     .get(`http://localhost:4200/digitalocean/account`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getActionsHistory(){
    return this.http
     .get(`${localAddress.address}/actions`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getContentDeliveryNetworks(){
    return this.http
     .get(`${localAddress.address}/cdns`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getCertificates(){
    return this.http
     .get(`${localAddress.address}/certificates`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getDomains(){
    return this.http
     .get(`${localAddress.address}/domains`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getDroplets(){
    return this.http
     .get(`${localAddress.address}/droplets`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getListOfFirewalls(){
    return this.http
     .get(`${localAddress.address}/firewalls/list`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getUnsecureFirewalls(){
    return this.http
     .get(`${localAddress.address}/firewalls/unsecure`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getFloatingIps(){
    return this.http
     .get(`${localAddress.address}/floatingips`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getKubernetesClusters(){
    return this.http
     .get(`${localAddress.address}/k8s`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getSshKeys(){
    return this.http
     .get(`${localAddress.address}/keys`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getLoadBalancers(){
    return this.http
     .get(`${localAddress.address}/loadbalancers`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getProjects(){
    return this.http
     .get(`${localAddress.address}/projects`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getRecords(){
    return this.http
     .get(`${localAddress.address}/records`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getSnapshots(){
    return this.http
     .get(`${localAddress.address}/snapshots`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getVolumes(){
    return this.http
     .get(`${localAddress.address}/volumes`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }

  public getDatabases(){
    return this.http
     .get(`${localAddress.address}/databases`)
     .map(res => {
       return res.json()
     })
     .catch(err => {
      let payload = JSON.parse(err._body)
      if (payload && payload.error)
        this.storeService.add(payload.error);
       return Observable.throw(err.json().error)
     })
  }


}
