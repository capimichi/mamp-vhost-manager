package controller

import (
	"context"
	"mamp-vhosts-manager/helper"
	"mamp-vhosts-manager/model"
)

type VhostController struct {
	ctx context.Context
}

func (v *VhostController) GetVhosts() []*model.Vhost {
	mampHelper := helper.MampHelper{}
	return mampHelper.GetVirtualHosts()
}

func (v *VhostController) GetVhost(name string) *model.Vhost {
	mampHelper := helper.MampHelper{}
	for _, vhost := range mampHelper.GetVirtualHosts(){
		if(vhost.Name == name){
			return vhost
		}
	}
	return nil
}

func (v *VhostController) GetGuessDocumentRoot() string {
	mampHelper := helper.MampHelper{}
	return mampHelper.GuessDocumentRoot()
}

func (v *VhostController) RestartApache(name string) {
	mampHelper := helper.MampHelper{}
	mampHelper.RestartApache()
}

// Function to create a new vhost
func (v *VhostController) CreateVhost(name string, serverName string, serverAdmin string, documentRoot string) *model.Vhost {
	mampHelper := helper.MampHelper{}
	return mampHelper.CreateVirtualHost(name, serverName, serverAdmin, documentRoot)
}

// Function to delete a vhost
func (v *VhostController) DeleteVhost(name string) {
	mampHelper := helper.MampHelper{}
	mampHelper.DeleteVirtualHost(name)
}