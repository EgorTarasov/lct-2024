/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as MapImport } from './routes/_map'
import { Route as BaseImport } from './routes/_base'
import { Route as MapHeatdistributorsImport } from './routes/_map/_heat_distributors'
import { Route as BaseIncidentsImport } from './routes/_base/_incidents'
import { Route as MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersImport } from './routes/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers'

// Create Virtual Routes

const BaseReportsLazyImport = createFileRoute('/_base/reports')()
const BaseRegisterLazyImport = createFileRoute('/_base/register')()
const BaseProfileLazyImport = createFileRoute('/_base/profile')()
const BaseLoginLazyImport = createFileRoute('/_base/login')()
const MapHeatdistributorsIndexLazyImport = createFileRoute(
  '/_map/_heat_distributors/',
)()
const MapHeatdistributorsHeatdistributorHeatDistributorIdImport =
  createFileRoute(
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId',
  )()
const BaseIncidentsIncidentsIndexLazyImport = createFileRoute(
  '/_base/_incidents/incidents/',
)()
const BaseIncidentsIncidentsUnomLazyImport = createFileRoute(
  '/_base/_incidents/incidents/$unom',
)()
const MapHeatdistributorsHeatdistributorHeatDistributorIdIndexLazyImport =
  createFileRoute(
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/',
  )()
const MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersIndexLazyImport =
  createFileRoute(
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/',
  )()
const MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersConsumerIdLazyImport =
  createFileRoute(
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId',
  )()

// Create/Update Routes

const MapRoute = MapImport.update({
  id: '/_map',
  getParentRoute: () => rootRoute,
} as any)

const BaseRoute = BaseImport.update({
  id: '/_base',
  getParentRoute: () => rootRoute,
} as any)

const BaseReportsLazyRoute = BaseReportsLazyImport.update({
  path: '/reports',
  getParentRoute: () => BaseRoute,
} as any).lazy(() => import('./routes/_base/reports.lazy').then((d) => d.Route))

const BaseRegisterLazyRoute = BaseRegisterLazyImport.update({
  path: '/register',
  getParentRoute: () => BaseRoute,
} as any).lazy(() =>
  import('./routes/_base/register.lazy').then((d) => d.Route),
)

const BaseProfileLazyRoute = BaseProfileLazyImport.update({
  path: '/profile',
  getParentRoute: () => BaseRoute,
} as any).lazy(() => import('./routes/_base/profile.lazy').then((d) => d.Route))

const BaseLoginLazyRoute = BaseLoginLazyImport.update({
  path: '/login',
  getParentRoute: () => BaseRoute,
} as any).lazy(() => import('./routes/_base/login.lazy').then((d) => d.Route))

const MapHeatdistributorsRoute = MapHeatdistributorsImport.update({
  id: '/_heat_distributors',
  getParentRoute: () => MapRoute,
} as any)

const BaseIncidentsRoute = BaseIncidentsImport.update({
  id: '/_incidents',
  getParentRoute: () => BaseRoute,
} as any)

const MapHeatdistributorsIndexLazyRoute =
  MapHeatdistributorsIndexLazyImport.update({
    path: '/',
    getParentRoute: () => MapHeatdistributorsRoute,
  } as any).lazy(() =>
    import('./routes/_map/_heat_distributors/index.lazy').then((d) => d.Route),
  )

const MapHeatdistributorsHeatdistributorHeatDistributorIdRoute =
  MapHeatdistributorsHeatdistributorHeatDistributorIdImport.update({
    path: '/heat_distributor/$heatDistributorId',
    getParentRoute: () => MapHeatdistributorsRoute,
  } as any)

const BaseIncidentsIncidentsIndexLazyRoute =
  BaseIncidentsIncidentsIndexLazyImport.update({
    path: '/incidents/',
    getParentRoute: () => BaseIncidentsRoute,
  } as any).lazy(() =>
    import('./routes/_base/_incidents/incidents/index.lazy').then(
      (d) => d.Route,
    ),
  )

const BaseIncidentsIncidentsUnomLazyRoute =
  BaseIncidentsIncidentsUnomLazyImport.update({
    path: '/incidents/$unom',
    getParentRoute: () => BaseIncidentsRoute,
  } as any).lazy(() =>
    import('./routes/_base/_incidents/incidents/$unom.lazy').then(
      (d) => d.Route,
    ),
  )

const MapHeatdistributorsHeatdistributorHeatDistributorIdIndexLazyRoute =
  MapHeatdistributorsHeatdistributorHeatDistributorIdIndexLazyImport.update({
    path: '/',
    getParentRoute: () =>
      MapHeatdistributorsHeatdistributorHeatDistributorIdRoute,
  } as any).lazy(() =>
    import(
      './routes/_map/_heat_distributors/heat_distributor/$heatDistributorId/index.lazy'
    ).then((d) => d.Route),
  )

const MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersRoute =
  MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersImport.update({
    id: '/_consumers',
    getParentRoute: () =>
      MapHeatdistributorsHeatdistributorHeatDistributorIdRoute,
  } as any)

const MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersIndexLazyRoute =
  MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersIndexLazyImport.update(
    {
      path: '/consumers/',
      getParentRoute: () =>
        MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersRoute,
    } as any,
  ).lazy(() =>
    import(
      './routes/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/index.lazy'
    ).then((d) => d.Route),
  )

const MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersConsumerIdLazyRoute =
  MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersConsumerIdLazyImport.update(
    {
      path: '/consumers/$consumerId',
      getParentRoute: () =>
        MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersRoute,
    } as any,
  ).lazy(() =>
    import(
      './routes/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId.lazy'
    ).then((d) => d.Route),
  )

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/_base': {
      id: '/_base'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof BaseImport
      parentRoute: typeof rootRoute
    }
    '/_map': {
      id: '/_map'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof MapImport
      parentRoute: typeof rootRoute
    }
    '/_base/_incidents': {
      id: '/_base/_incidents'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof BaseIncidentsImport
      parentRoute: typeof BaseImport
    }
    '/_map/_heat_distributors': {
      id: '/_map/_heat_distributors'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof MapHeatdistributorsImport
      parentRoute: typeof MapImport
    }
    '/_base/login': {
      id: '/_base/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof BaseLoginLazyImport
      parentRoute: typeof BaseImport
    }
    '/_base/profile': {
      id: '/_base/profile'
      path: '/profile'
      fullPath: '/profile'
      preLoaderRoute: typeof BaseProfileLazyImport
      parentRoute: typeof BaseImport
    }
    '/_base/register': {
      id: '/_base/register'
      path: '/register'
      fullPath: '/register'
      preLoaderRoute: typeof BaseRegisterLazyImport
      parentRoute: typeof BaseImport
    }
    '/_base/reports': {
      id: '/_base/reports'
      path: '/reports'
      fullPath: '/reports'
      preLoaderRoute: typeof BaseReportsLazyImport
      parentRoute: typeof BaseImport
    }
    '/_map/_heat_distributors/': {
      id: '/_map/_heat_distributors/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof MapHeatdistributorsIndexLazyImport
      parentRoute: typeof MapHeatdistributorsImport
    }
    '/_base/_incidents/incidents/$unom': {
      id: '/_base/_incidents/incidents/$unom'
      path: '/incidents/$unom'
      fullPath: '/incidents/$unom'
      preLoaderRoute: typeof BaseIncidentsIncidentsUnomLazyImport
      parentRoute: typeof BaseIncidentsImport
    }
    '/_base/_incidents/incidents/': {
      id: '/_base/_incidents/incidents/'
      path: '/incidents'
      fullPath: '/incidents'
      preLoaderRoute: typeof BaseIncidentsIncidentsIndexLazyImport
      parentRoute: typeof BaseIncidentsImport
    }
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId': {
      id: '/_map/_heat_distributors/heat_distributor/$heatDistributorId'
      path: '/heat_distributor/$heatDistributorId'
      fullPath: '/heat_distributor/$heatDistributorId'
      preLoaderRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdImport
      parentRoute: typeof MapHeatdistributorsImport
    }
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers': {
      id: '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers'
      path: '/heat_distributor/$heatDistributorId'
      fullPath: '/heat_distributor/$heatDistributorId'
      preLoaderRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersImport
      parentRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdRoute
    }
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/': {
      id: '/_map/_heat_distributors/heat_distributor/$heatDistributorId/'
      path: '/'
      fullPath: '/heat_distributor/$heatDistributorId/'
      preLoaderRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdIndexLazyImport
      parentRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdImport
    }
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId': {
      id: '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId'
      path: '/consumers/$consumerId'
      fullPath: '/heat_distributor/$heatDistributorId/consumers/$consumerId'
      preLoaderRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersConsumerIdLazyImport
      parentRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersImport
    }
    '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/': {
      id: '/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/'
      path: '/consumers'
      fullPath: '/heat_distributor/$heatDistributorId/consumers'
      preLoaderRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersIndexLazyImport
      parentRoute: typeof MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren({
  BaseRoute: BaseRoute.addChildren({
    BaseIncidentsRoute: BaseIncidentsRoute.addChildren({
      BaseIncidentsIncidentsUnomLazyRoute,
      BaseIncidentsIncidentsIndexLazyRoute,
    }),
    BaseLoginLazyRoute,
    BaseProfileLazyRoute,
    BaseRegisterLazyRoute,
    BaseReportsLazyRoute,
  }),
  MapRoute: MapRoute.addChildren({
    MapHeatdistributorsRoute: MapHeatdistributorsRoute.addChildren({
      MapHeatdistributorsIndexLazyRoute,
      MapHeatdistributorsHeatdistributorHeatDistributorIdRoute:
        MapHeatdistributorsHeatdistributorHeatDistributorIdRoute.addChildren({
          MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersRoute:
            MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersRoute.addChildren(
              {
                MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersConsumerIdLazyRoute,
                MapHeatdistributorsHeatdistributorHeatDistributorIdConsumersConsumersIndexLazyRoute,
              },
            ),
          MapHeatdistributorsHeatdistributorHeatDistributorIdIndexLazyRoute,
        }),
    }),
  }),
})

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/_base",
        "/_map"
      ]
    },
    "/_base": {
      "filePath": "_base.tsx",
      "children": [
        "/_base/_incidents",
        "/_base/login",
        "/_base/profile",
        "/_base/register",
        "/_base/reports"
      ]
    },
    "/_map": {
      "filePath": "_map.tsx",
      "children": [
        "/_map/_heat_distributors"
      ]
    },
    "/_base/_incidents": {
      "filePath": "_base/_incidents.tsx",
      "parent": "/_base",
      "children": [
        "/_base/_incidents/incidents/$unom",
        "/_base/_incidents/incidents/"
      ]
    },
    "/_map/_heat_distributors": {
      "filePath": "_map/_heat_distributors.tsx",
      "parent": "/_map",
      "children": [
        "/_map/_heat_distributors/",
        "/_map/_heat_distributors/heat_distributor/$heatDistributorId"
      ]
    },
    "/_base/login": {
      "filePath": "_base/login.lazy.tsx",
      "parent": "/_base"
    },
    "/_base/profile": {
      "filePath": "_base/profile.lazy.tsx",
      "parent": "/_base"
    },
    "/_base/register": {
      "filePath": "_base/register.lazy.tsx",
      "parent": "/_base"
    },
    "/_base/reports": {
      "filePath": "_base/reports.lazy.tsx",
      "parent": "/_base"
    },
    "/_map/_heat_distributors/": {
      "filePath": "_map/_heat_distributors/index.lazy.tsx",
      "parent": "/_map/_heat_distributors"
    },
    "/_base/_incidents/incidents/$unom": {
      "filePath": "_base/_incidents/incidents/$unom.lazy.tsx",
      "parent": "/_base/_incidents"
    },
    "/_base/_incidents/incidents/": {
      "filePath": "_base/_incidents/incidents/index.lazy.tsx",
      "parent": "/_base/_incidents"
    },
    "/_map/_heat_distributors/heat_distributor/$heatDistributorId": {
      "filePath": "_map/_heat_distributors/heat_distributor/$heatDistributorId",
      "parent": "/_map/_heat_distributors",
      "children": [
        "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers",
        "/_map/_heat_distributors/heat_distributor/$heatDistributorId/"
      ]
    },
    "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers": {
      "filePath": "_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers.tsx",
      "parent": "/_map/_heat_distributors/heat_distributor/$heatDistributorId",
      "children": [
        "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId",
        "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/"
      ]
    },
    "/_map/_heat_distributors/heat_distributor/$heatDistributorId/": {
      "filePath": "_map/_heat_distributors/heat_distributor/$heatDistributorId/index.lazy.tsx",
      "parent": "/_map/_heat_distributors/heat_distributor/$heatDistributorId"
    },
    "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId": {
      "filePath": "_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/$consumerId.lazy.tsx",
      "parent": "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers"
    },
    "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/": {
      "filePath": "_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers/consumers/index.lazy.tsx",
      "parent": "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers"
    }
  }
}
ROUTE_MANIFEST_END */
