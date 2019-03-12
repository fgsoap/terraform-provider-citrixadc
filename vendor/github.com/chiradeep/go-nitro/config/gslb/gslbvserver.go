package gslb

type Gslbvserver struct {
	Activeservices            int    `json:"activeservices,omitempty"`
	Appflowlog                string `json:"appflowlog,omitempty"`
	Backupip                  string `json:"backupip,omitempty"`
	Backuplbmethod            string `json:"backuplbmethod,omitempty"`
	Backupsessiontimeout      int    `json:"backupsessiontimeout,omitempty"`
	Backupvserver             string `json:"backupvserver,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	Considereffectivestate    string `json:"considereffectivestate,omitempty"`
	Cookiedomain              string `json:"cookie_domain,omitempty"`
	Cookietimeout             int    `json:"cookietimeout,omitempty"`
	Curstate                  string `json:"curstate,omitempty"`
	Disableprimaryondown      string `json:"disableprimaryondown,omitempty"`
	Dnsrecordtype             string `json:"dnsrecordtype,omitempty"`
	Domainname                string `json:"domainname,omitempty"`
	Dynamicweight             string `json:"dynamicweight,omitempty"`
	Ecs                       string `json:"ecs,omitempty"`
	Ecsaddrvalidation         string `json:"ecsaddrvalidation,omitempty"`
	Edr                       string `json:"edr,omitempty"`
	Gotopriorityexpression    string `json:"gotopriorityexpression,omitempty"`
	Health                    int    `json:"health,omitempty"`
	Iptype                    string `json:"iptype,omitempty"`
	Iscname                   string `json:"iscname,omitempty"`
	Lbmethod                  string `json:"lbmethod,omitempty"`
	Lbrrreason                int    `json:"lbrrreason,omitempty"`
	Mir                       string `json:"mir,omitempty"`
	Name                      string `json:"name,omitempty"`
	Netmask                   string `json:"netmask,omitempty"`
	Newname                   string `json:"newname,omitempty"`
	Nodefaultbindings         string `json:"nodefaultbindings,omitempty"`
	Persistenceid             int    `json:"persistenceid,omitempty"`
	Persistencetype           string `json:"persistencetype,omitempty"`
	Persistmask               string `json:"persistmask,omitempty"`
	Policyname                string `json:"policyname,omitempty"`
	Priority                  int    `json:"priority,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicename               string `json:"servicename,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
	Sitedomainttl             int    `json:"sitedomainttl,omitempty"`
	Sitepersistence           string `json:"sitepersistence,omitempty"`
	Sobackupaction            string `json:"sobackupaction,omitempty"`
	Somethod                  string `json:"somethod,omitempty"`
	Sopersistence             string `json:"sopersistence,omitempty"`
	Sopersistencetimeout      int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold               int    `json:"sothreshold,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimemsec       int    `json:"statechangetimemsec,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Status                    int    `json:"status,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Timeout                   int    `json:"timeout,omitempty"`
	Tolerance                 int    `json:"tolerance,omitempty"`
	Totalservices             int    `json:"totalservices,omitempty"`
	Ttl                       int    `json:"ttl,omitempty"`
	Type                      string `json:"type,omitempty"`
	V6netmasklen              int    `json:"v6netmasklen,omitempty"`
	V6persistmasklen          int    `json:"v6persistmasklen,omitempty"`
	Vsvrbindsvcip             string `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport           int    `json:"vsvrbindsvcport,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}
