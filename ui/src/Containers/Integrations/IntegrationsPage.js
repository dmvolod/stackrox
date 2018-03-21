import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { ToastContainer, toast } from 'react-toastify';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';

import IntegrationModal from 'Containers/Integrations/IntegrationModal';
import IntegrationTile from 'Containers/Integrations/IntegrationTile';
import { actions as authActions } from 'reducers/auth';
import { actions } from 'reducers/integrations';
import { selectors } from 'reducers';

import auth0 from 'images/auth0.svg';
import docker from 'images/docker.svg';
import jira from 'images/jira.svg';
import kubernetes from 'images/kubernetes.svg';
import slack from 'images/slack.svg';
import tenable from 'images/tenable.svg';
import email from 'images/email.svg';
import quay from 'images/quay.svg';
import clair from 'images/clair.svg';

const dataSources = {
    authProviders: [
        {
            label: 'Auth0',
            type: 'auth0',
            source: 'authProviders',
            image: auth0
        }
    ],
    registries: [
        {
            label: 'Docker Registry',
            type: 'docker',
            source: 'registries',
            image: docker
        },
        {
            label: 'Tenable Registry',
            type: 'tenable',
            source: 'registries',
            image: tenable
        },
        {
            label: 'Quay Registry',
            type: 'quay',
            source: 'registries',
            image: quay
        }
    ],
    orchestrators: [
        {
            label: 'Docker Enterprise Edition',
            image: docker,
            source: 'clusters',
            type: 'DOCKER_EE_CLUSTER'
        },
        {
            label: 'Kubernetes',
            image: kubernetes,
            source: 'clusters',
            type: 'KUBERNETES_CLUSTER'
        },
        {
            label: 'Docker Swarm',
            image: docker,
            source: 'clusters',
            type: 'SWARM_CLUSTER'
        }
    ],
    scanners: [
        {
            label: 'Docker Trusted Registry',
            type: 'dtr',
            source: 'scanners',
            image: docker
        },
        {
            label: 'Tenable',
            type: 'tenable',
            source: 'scanners',
            image: tenable
        },
        {
            label: 'Quay',
            type: 'quay',
            source: 'scanners',
            image: quay
        },
        {
            label: 'Clair',
            type: 'clair',
            source: 'scanners',
            image: clair
        }
    ],
    plugins: [
        {
            label: 'Slack',
            type: 'slack',
            source: 'notifiers',
            image: slack
        },
        {
            label: 'Jira',
            type: 'jira',
            source: 'notifiers',
            image: jira
        },
        {
            label: 'Email',
            type: 'email',
            source: 'notifiers',
            image: email
        }
    ]
};

const reducer = (action, prevState, nextState) => {
    switch (action) {
        case 'OPEN_INTEGRATION_MODAL':
            return {
                integrationModal: {
                    open: true,
                    integrations: nextState.integrations,
                    source: nextState.source,
                    type: nextState.type
                }
            };
        case 'CLOSE_INTEGRATION_MODAL':
            return {
                integrationModal: {
                    open: false,
                    integrations: [],
                    source: '',
                    type: ''
                }
            };
        default:
            return prevState;
    }
};

class IntegrationsPage extends Component {
    static propTypes = {
        /* eslint-disable */
        authProviders: PropTypes.arrayOf(
            PropTypes.shape({
                config: PropTypes.shape({
                    audience: PropTypes.string.isRequired,
                    client_id: PropTypes.string.isRequired,
                    domain: PropTypes.string.isRequired
                }),
                name: PropTypes.string.isRequired
            })
        ).isRequired,
        clusters: PropTypes.arrayOf(PropTypes.object).isRequired,
        notifiers: PropTypes.arrayOf(PropTypes.object).isRequired,
        registries: PropTypes.arrayOf(PropTypes.object).isRequired,
        scanners: PropTypes.arrayOf(PropTypes.object).isRequired,
        /* eslint-enable */
        fetchAuthProviders: PropTypes.func.isRequired,
        fetchNotifiers: PropTypes.func.isRequired,
        fetchRegistries: PropTypes.func.isRequired,
        fetchScanners: PropTypes.func.isRequired
    };

    constructor(props) {
        super(props);

        this.state = {
            integrationModal: {
                open: false,
                integrations: [],
                source: '',
                type: ''
            }
        };
    }

    getEntities = source => {
        switch (source) {
            case 'authProviders':
                this.props.fetchAuthProviders();
                break;
            case 'scanners':
                this.props.fetchScanners();
                break;
            case 'registries':
                this.props.fetchRegistries();
                break;
            case 'notifiers':
                this.props.fetchNotifiers();
                break;
            default:
                throw new Error(`Unknown source ${source}`);
        }
    };

    getClustersForOrchestrator = orchestrator => {
        const { type } = orchestrator;
        const clusters = this.props.clusters.filter(cluster => cluster.type === type);
        return clusters;
    };

    openIntegrationModal = integrationCategory => {
        const { source, type } = integrationCategory;
        const integrations =
            source !== 'clusters'
                ? this.props[source].filter(i => i.type === type.toLowerCase())
                : this.props.clusters.filter(cluster => cluster.type === type);
        this.update('OPEN_INTEGRATION_MODAL', { integrations, source, type });
    };

    closeIntegrationModal = isSuccessful => {
        if (isSuccessful === true) {
            const { integrationModal: { source, type } } = this.state;
            toast(`Successfully integrated ${type}`);
            this.getEntities(source);
        }
        this.update('CLOSE_INTEGRATION_MODAL');
    };

    findIntegrations = (source, type) => {
        const integrations = this.props[source].filter(i => i.type === type.toLowerCase());
        return integrations.filter(obj => obj.type === type);
    };

    update = (action, nextState) => {
        this.setState(prevState => reducer(action, prevState, nextState));
    };

    renderIntegrationModal() {
        const { integrationModal: { source, type, open } } = this.state;
        if (!open) return null;
        const integrations =
            source !== 'clusters'
                ? this.props[source].filter(i => i.type === type.toLowerCase())
                : this.props.clusters.filter(cluster => cluster.type === type);
        return (
            <IntegrationModal
                integrations={integrations}
                source={source}
                type={type}
                onRequestClose={this.closeIntegrationModal}
                onIntegrationsUpdate={this.getEntities}
            />
        );
    }

    renderIntegrationTiles = source =>
        dataSources[source].map(tile => (
            <IntegrationTile
                key={tile.label}
                integration={tile}
                onClick={this.openIntegrationModal}
                numIntegrations={
                    source !== 'orchestrators'
                        ? this.findIntegrations(tile.source, tile.type).length
                        : this.getClustersForOrchestrator(tile).length
                }
            />
        ));

    render() {
        const registries = this.renderIntegrationTiles('registries');
        const orchestrators = this.renderIntegrationTiles('orchestrators');
        const scanners = this.renderIntegrationTiles('scanners');
        const plugins = this.renderIntegrationTiles('plugins');
        const authProviders = this.renderIntegrationTiles('authProviders');

        return (
            <section className="flex">
                <ToastContainer
                    toastClassName="font-sans text-base-600 text-white font-600 bg-black"
                    hideProgressBar
                    autoClose={3000}
                />
                <div className="md:w-full border-r border-primary-300 pt-4">
                    <h1 className="font-500 mx-3 border-b border-primary-300 pb-4 uppercase text-xl font-800 text-primary-600 tracking-wide">
                        Data sources
                    </h1>
                    <div>
                        <h2 className="mx-3 mt-8 text-xl text-base text-primary-500 pb-3">
                            Registries
                        </h2>
                        <div className="flex flex-wrap">{registries}</div>
                    </div>
                    <div>
                        <h2 className="mx-3 mt-8 text-xl text-base text-primary-500 border-t border-primary-300 pt-6 pb-3">
                            Orchestrators &amp; Container Platforms
                        </h2>
                        <div className="flex flex-wrap">{orchestrators}</div>
                    </div>
                    <div className="mb-6">
                        <h2 className="mx-3 mt-8 text-xl text-base text-primary-500 border-t border-primary-300 pt-6 pb-3">
                            Scanning &amp; Governance Tools
                        </h2>
                        <div className="flex flex-wrap">{scanners}</div>
                    </div>
                    <div className="mb-6">
                        <h2 className="mx-3 mt-8 text-xl text-base text-primary-500 border-t border-primary-300 pt-6 pb-3">
                            Plugins
                        </h2>
                        <div className="flex flex-wrap">{plugins}</div>
                    </div>
                    <div className="mb-6">
                        <h2 className="mx-3 mt-8 text-xl text-base text-primary-500 border-t border-primary-300 pt-6 pb-3">
                            Authentication Providers
                        </h2>
                        <div className="flex flex-wrap">{authProviders}</div>
                    </div>
                </div>
                {this.renderIntegrationModal()}
            </section>
        );
    }
}

const mapStateToProps = createStructuredSelector({
    authProviders: selectors.getAuthProviders,
    clusters: selectors.getClusters,
    notifiers: selectors.getNotifiers,
    registries: selectors.getRegistries,
    scanners: selectors.getScanners
});

const mapDispatchToProps = dispatch => ({
    fetchAuthProviders: () => dispatch(authActions.fetchAuthProviders.request()),
    fetchNotifiers: () => dispatch(actions.fetchNotifiers.request()),
    fetchRegistries: () => dispatch(actions.fetchRegistries.request()),
    fetchScanners: () => dispatch(actions.fetchScanners.request())
});

export default connect(mapStateToProps, mapDispatchToProps)(IntegrationsPage);
